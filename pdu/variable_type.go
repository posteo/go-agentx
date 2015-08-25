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

import "fmt"

// The various variable types.
const (
	VariableTypeInteger          VariableType = 2
	VariableTypeOctetString      VariableType = 4
	VariableTypeNull             VariableType = 5
	VariableTypeObjectIdentifier VariableType = 6
	VariableTypeIPAddress        VariableType = 64
	VariableTypeCounter32        VariableType = 65
	VariableTypeGauge32          VariableType = 66
	VariableTypeTimeTicks        VariableType = 67
	VariableTypeOpaque           VariableType = 68
	VariableTypeCounter64        VariableType = 70
	VariableTypeNoSuchObject     VariableType = 128
	VariableTypeNoSuchInstance   VariableType = 129
	VariableTypeEndOfMIBView     VariableType = 130
)

// VariableType defines the type of a variable.
type VariableType uint16

func (v VariableType) String() string {
	switch v {
	case VariableTypeInteger:
		return "VariableTypeInteger"
	case VariableTypeOctetString:
		return "VariableTypeOctetString"
	case VariableTypeNull:
		return "VariableTypeNull"
	case VariableTypeObjectIdentifier:
		return "VariableTypeObjectIdentifier"
	case VariableTypeIPAddress:
		return "VariableTypeIPAddress"
	case VariableTypeCounter32:
		return "VariableTypeCounter32"
	case VariableTypeGauge32:
		return "VariableTypeGauge32"
	case VariableTypeTimeTicks:
		return "VariableTypeTimeTicks"
	case VariableTypeOpaque:
		return "VariableTypeOpaque"
	case VariableTypeCounter64:
		return "VariableTypeCounter64"
	case VariableTypeNoSuchObject:
		return "VariableTypeNoSuchObject"
	case VariableTypeNoSuchInstance:
		return "VariableTypeNoSuchInstance"
	case VariableTypeEndOfMIBView:
		return "VariableTypeEndOfMIBView"
	}
	return fmt.Sprintf("VariableTypeUnknown (%d)", v)
}
