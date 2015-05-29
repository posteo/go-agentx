package agentx

import (
	"bufio"
	"log"
	"net"
	"time"

	"github.com/juju/errgo"
	"github.com/posteo/go-agentx/pdu"
)

// Client defines an agentx client.
type Client struct {
	Net     string
	Address string
	Timeout time.Duration
	NameOID string
	Name    string
	Debug   bool

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

	writer := bufio.NewWriter(c.connection)
	tx := make(chan *pdu.HeaderPacket)
	go func() {
		for headerPacket := range tx {
			if c.Debug {
				log.Printf("out %v", headerPacket)
			}
			headerPacketBytes, err := headerPacket.MarshalBinary()
			if err != nil {
				panic(err)
			}
			if _, err := writer.Write(headerPacketBytes); err != nil {
				panic(err)
			}
			if err := writer.Flush(); err != nil {
				panic(err)
			}
			if c.Debug {
				log.Printf("sent (%2d) %x", len(headerPacketBytes), headerPacketBytes)
			}
		}
	}()

	reader := bufio.NewReader(c.connection)
	rx := make(chan *pdu.HeaderPacket)
	go func() {
		for {
			headerBytes := make([]byte, pdu.HeaderSize)
			if _, err := reader.Read(headerBytes); err != nil {
				panic(err)
			}
			if c.Debug {
				log.Printf("recv (%2d) %x", len(headerBytes), headerBytes)
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
			default:
				log.Printf("unhandled packet of type %s", header.Type)
				continue
			}

			packetBytes := make([]byte, header.PayloadLength)
			if _, err := reader.Read(packetBytes); err != nil {
				panic(err)
			}
			if c.Debug {
				log.Printf("recv (%2d) %x", len(packetBytes), packetBytes)
			}

			if err := packet.UnmarshalBinary(packetBytes); err != nil {
				panic(err)
			}

			headerPacket := &pdu.HeaderPacket{Header: header, Packet: packet}
			if c.Debug {
				log.Printf("in %v", headerPacket)
			}

			rx <- headerPacket
		}
	}()

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
