// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import (
	"context"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

// Handler defines an interface for a handler of events that
// might occure during a session.
type Handler interface {
	Get(context.Context, value.OID) (value.OID, pdu.VariableType, any, error)
	GetNext(context.Context, value.OID, bool, value.OID) (value.OID, pdu.VariableType, any, error)
}

type (
	sessionIDKey     struct{}
	transactionIDKey struct{}
	packetIDKey      struct{}
)

func SessionID(ctx context.Context) uint32 {
	value, _ := ctx.Value(sessionIDKey{}).(uint32)
	return value
}

func withSessionID(ctx context.Context, value uint32) context.Context {
	return context.WithValue(ctx, sessionIDKey{}, value)
}

func TransactionID(ctx context.Context) uint32 {
	value, _ := ctx.Value(transactionIDKey{}).(uint32)
	return value
}

func withTransactionID(ctx context.Context, value uint32) context.Context {
	return context.WithValue(ctx, transactionIDKey{}, value)
}

func PacketID(ctx context.Context) uint32 {
	value, _ := ctx.Value(packetIDKey{}).(uint32)
	return value
}

func withPacketID(ctx context.Context, value uint32) context.Context {
	return context.WithValue(ctx, packetIDKey{}, value)
}
