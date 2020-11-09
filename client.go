// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

// Client defines an agentx client.
type Client struct {
	Timeout           time.Duration
	ReconnectInterval time.Duration
	NameOID           value.OID
	Name              string

	network     string
	address     string
	conn        net.Conn
	requestChan chan *request
	sessions    map[uint32]*Session
}

// Dial connects to the provided agentX endpoint.
func Dial(network, address string) (*Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, fmt.Errorf("dial %s %s: %w", network, address, err)
	}
	c := &Client{
		network:     network,
		address:     address,
		conn:        conn,
		requestChan: make(chan *request),
		sessions:    make(map[uint32]*Session),
	}
	tx := c.runTransmitter()
	rx := c.runReceiver()
	c.runDispatcher(tx, rx)

	return c, nil
}

// Close tears down the client.
func (c *Client) Close() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("close connection: %w", err)
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
		return nil, err
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
				log.Printf("marshal error: %v", err)
				continue
			}
			writer := bufio.NewWriter(c.conn)
			if _, err := writer.Write(headerPacketBytes); err != nil {
				log.Printf("write error: %v", err)
				continue
			}
			if err := writer.Flush(); err != nil {
				log.Printf("flush error: %v", err)
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
			reader := bufio.NewReader(c.conn)
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
						conn, err := net.Dial(c.network, c.address)
						if err != nil {
							log.Printf("try to reconnect: %v", err)
							continue reopenLoop
						}
						c.conn = conn
						go func() {
							for _, session := range c.sessions {
								delete(c.sessions, session.ID())
								if err := session.reopen(); err != nil {
									log.Printf("error during reopen session: %v", err)
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
	go func() {
		currentPacketID := uint32(0)
		responseChans := make(map[uint32]chan *pdu.HeaderPacket)

		for {
			select {
			case request := <-c.requestChan:
				// log.Printf(">: %v", request)
				request.headerPacket.Header.PacketID = currentPacketID
				responseChans[currentPacketID] = request.responseChan
				currentPacketID++

				tx <- request.headerPacket
			case headerPacket := <-rx:
				// log.Printf("<: %v", headerPacket)
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
