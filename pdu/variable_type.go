// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

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
