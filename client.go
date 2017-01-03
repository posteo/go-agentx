/*
go-agentx
Copyright (C) 2015 Philipp Brüll <bruell@simia.tech>

This library is free software; you can redistribute it and/or
modify it under the terms of the GNU Lesser General Public
License as published by the Free Software Foundation; either
version 2.1 of the License, or (at your option) any later version.

This library is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public
License along with this library; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301
USA
*/

package agentx

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/martinclaro/go-agentx/pdu"
	"github.com/martinclaro/go-agentx/value"
	"gopkg.in/errgo.v1"
)

// Client defines an agentx client.
type Client struct {
	Net               string
	Address           string
	Timeout           time.Duration
	ReconnectInterval time.Duration
	NameOID           value.OID
	Name              string

	connection  net.Conn
	requestChan chan *request
	sessions    map[uint32]*Session
}

// Open sets up the client.
func (c *Client) Open() error {
	connection, err := net.Dial(c.Net, c.Address)
	if err != nil {
		return errgo.Mask(err)
	}
	c.connection = connection
	c.sessions = make(map[uint32]*Session)

	tx := c.runTransmitter()
	rx := c.runReceiver()
	c.runDispatcher(tx, rx)

	return nil
}

// Close tears down the client.
func (c *Client) Close() error {
	if err := c.connection.Close(); err != nil {
		return errgo.Mask(err)
	}
	return nil
}

// Session sets up a new session.
func (c *Client) Session() (*Session, error) {
	s := &Session{
		client:  c,
		timeout: c.Timeout,
	}
	if err := s.open(c.NameOID, c.Name); err != nil {
		return nil, errgo.Mask(err)
	}
	c.sessions[s.ID()] = s

	return s, nil
}

func (c *Client) runTransmitter() chan *pdu.HeaderPacket {
	tx := make(chan *pdu.HeaderPacket)

	go func() {
		for headerPacket := range tx {
			headerPacketBytes, err := headerPacket.MarshalBinary()
			if err != nil {
				log.Printf(errgo.Details(err))
				continue
			}
			writer := bufio.NewWriter(c.connection)
			if _, err := writer.Write(headerPacketBytes); err != nil {
				log.Printf(errgo.Details(err))
				continue
			}
			if err := writer.Flush(); err != nil {
				log.Printf(errgo.Details(err))
				continue
			}
		}
	}()

	return tx
}

func (c *Client) runReceiver() chan *pdu.HeaderPacket {
	rx := make(chan *pdu.HeaderPacket)

	go func() {
	mainLoop:
		for {
			reader := bufio.NewReader(c.connection)
			headerBytes := make([]byte, pdu.HeaderSize)
			if _, err := reader.Read(headerBytes); err != nil {
				if opErr, ok := err.(*net.OpError); ok && strings.HasSuffix(opErr.Error(), "use of closed network connection") {
					return
				}
				if err == io.EOF {
					log.Printf("lost connection - try to re-connect ...")
				reopenLoop:
					for {
						time.Sleep(c.ReconnectInterval)
						connection, err := net.Dial(c.Net, c.Address)
						if err != nil {
							log.Printf("try to reconnect: %s", errgo.Details(err))
							continue reopenLoop
						}
						c.connection = connection
						go func() {
							for _, session := range c.sessions {
								delete(c.sessions, session.ID())
								if err := session.reopen(); err != nil {
									log.Printf("error during reopen session: %s", errgo.Details(err))
									return
								}
								c.sessions[session.ID()] = session
								log.Printf("successful re-connected")
							}
						}()
						continue mainLoop
					}
				}
				panic(err)
			}

			header := &pdu.Header{}
			if err := header.UnmarshalBinary(headerBytes); err != nil {
				panic(err)
			}

			var packet pdu.Packet
			switch header.Type {
			case pdu.TypeResponse:
				packet = &pdu.Response{}
			case pdu.TypeGet:
				packet = &pdu.Get{}
			case pdu.TypeGetNext:
				packet = &pdu.GetNext{}
			default:
				log.Printf("unhandled packet of type %s", header.Type)
			}

			packetBytes := make([]byte, header.PayloadLength)
			if _, err := reader.Read(packetBytes); err != nil {
				panic(err)
			}

			if err := packet.UnmarshalBinary(packetBytes); err != nil {
				panic(err)
			}

			rx <- &pdu.HeaderPacket{Header: header, Packet: packet}
		}
	}()

	return rx
}

func (c *Client) runDispatcher(tx, rx chan *pdu.HeaderPacket) {
	c.requestChan = make(chan *request)

	go func() {
		currentPacketID := uint32(0)
		responseChans := make(map[uint32]chan *pdu.HeaderPacket)

		for {
			select {
			case request := <-c.requestChan:
				request.headerPacket.Header.PacketID = currentPacketID
				responseChans[currentPacketID] = request.responseChan
				currentPacketID++

				tx <- request.headerPacket
			case headerPacket := <-rx:
				packetID := headerPacket.Header.PacketID
				responseChan, ok := responseChans[packetID]
				if ok {
					responseChan <- headerPacket
					delete(responseChans, packetID)
				} else {
					session, ok := c.sessions[headerPacket.Header.SessionID]
					if ok {
						tx <- session.handle(headerPacket)
					} else {
						log.Printf("got without session: %v", headerPacket)
					}
				}
			}
		}
	}()
}

func (c *Client) request(hp *pdu.HeaderPacket) *pdu.HeaderPacket {
	responseChan := make(chan *pdu.HeaderPacket)
	request := &request{
		headerPacket: hp,
		responseChan: responseChan,
	}
	c.requestChan <- request
	headerPacket := <-responseChan
	return headerPacket
}
