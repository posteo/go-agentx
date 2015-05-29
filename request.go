package agentx

import "github.com/posteo/go-agentx/pdu"

type request struct {
	headerPacket *pdu.HeaderPacket
	responseChan chan *pdu.HeaderPacket
}
