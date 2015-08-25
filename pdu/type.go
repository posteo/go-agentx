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
