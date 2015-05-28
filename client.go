package agentx

import (
	"bufio"
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
	RootOID string
	Name    string

	connection net.Conn
	readWriter *bufio.ReadWriter
}

// Open sets up the client.
func (c *Client) Open() error {
	connection, err := net.Dial(c.Net, c.Address)
	if err != nil {
		return errgo.Mask(err)
	}
	c.connection = connection
	c.readWriter = bufio.NewReadWriter(bufio.NewReader(c.connection), bufio.NewWriter(c.connection))

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
	request := &pdu.Open{}
	request.Timeout.Duration = c.Timeout
	request.ID.SetByOID(c.RootOID)
	request.Description.Text = c.Name

	if err := WritePacket(c.readWriter, request); err != nil {
		return nil, errgo.Mask(err)
	}
	if err := c.readWriter.Flush(); err != nil {
		return nil, errgo.Mask(err)
	}

	header, _, err := ReadPacket(c.readWriter)
	if err != nil {
		return nil, errgo.Mask(err)
	}

	return &Session{
		readWriter: c.readWriter,
		sessionID:  header.SessionID,
	}, nil
}
