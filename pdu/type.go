// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package pdu

// The various pdu packet types.
const (
	TypeOpen            Type = 1
	TypeClose           Type = 2
	TypeRegister        Type = 3
	TypeUnregister      Type = 4
	TypeGet             Type = 5
	TypeGetNext         Type = 6
	TypeGetBulk         Type = 7
	TypeTestSet         Type = 8
	TypeCommitSet       Type = 9
	TypeUndoSet         Type = 10
	TypeCleanupSet      Type = 11
	TypeNotify          Type = 12
	TypePing            Type = 13
	TypeIndexAllocate   Type = 14
	TypeIndexDeallocate Type = 15
	TypeAddAgentCaps    Type = 16
	TypeRemoveAgentCaps Type = 17
	TypeResponse        Type = 18
)

// Type defines the pdu packet type.
type Type byte

// TypeOwner defines the interface for an object that provides a type.
type TypeOwner interface {
	Type() Type
}

func (t Type) String() string {
	switch t {
	case TypeOpen:
		return "TypeOpen"
	case TypeClose:
		return "TypeClose"
	case TypeRegister:
		return "TypeRegister"
	case TypeUnregister:
		return "TypeUnregister"
	case TypeGet:
		return "TypeGet"
	case TypeGetNext:
		return "TypeGetNext"
	case TypeGetBulk:
		return "TypeGetBulk"
	case TypeTestSet:
		return "TypeTestSet"
	case TypeCommitSet:
		return "TypeCommitSet"
	case TypeUndoSet:
		return "TypeUndoSet"
	case TypeCleanupSet:
		return "TypeCleanupSet"
	case TypeNotify:
		return "TypeNotify"
	case TypePing:
		return "TypePing"
	case TypeIndexAllocate:
		return "TypeIndexAllocate"
	case TypeIndexDeallocate:
		return "TypeIndexDeallocate"
	case TypeAddAgentCaps:
		return "TypeAddAgentCaps"
	case TypeRemoveAgentCaps:
		return "TypeRemoveAgentCaps"
	case TypeResponse:
		return "TypeResponse"
	}
	return "TypeUnknown"
}
