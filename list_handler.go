// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import (
	"bytes"
	"sort"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

// ListHandler is a helper that takes a list of oids and implements
// a default behaviour for that list.
type ListHandler struct {
	oids  sort.StringSlice
	items map[string]*ListItem
}

// Add adds a list item for the provided oid and returns it.
func (l *ListHandler) Add(oid string) *ListItem {
	if l.items == nil {
		l.items = make(map[string]*ListItem)
	}

	l.oids = append(l.oids, oid)
	l.oids.Sort()
	item := &ListItem{}
	l.items[oid] = item
	return item
}

// Get tries to find the provided oid and returns the corresponding value.
func (l *ListHandler) Get(oid value.OID) (value.OID, pdu.VariableType, interface{}, error) {
	if l.items == nil {
		return nil, pdu.VariableTypeNoSuchObject, nil, nil
	}

	item, ok := l.items[oid.String()]
	if ok {
		return oid, item.Type, item.Value, nil
	}
	return nil, pdu.VariableTypeNoSuchObject, nil, nil
}

// GetNext tries to find the value that follows the provided oid and returns it.
func (l *ListHandler) GetNext(from value.OID, includeFrom bool, to value.OID) (value.OID, pdu.VariableType, interface{}, error) {
	if l.items == nil {
		return nil, pdu.VariableTypeNoSuchObject, nil, nil
	}

	fromOID, toOID := from.String(), to.String()
	for _, oid := range l.oids {
		if oidWithin(oid, fromOID, includeFrom, toOID) {
			return l.Get(value.MustParseOID(oid))
		}
	}

	return nil, pdu.VariableTypeNoSuchObject, nil, nil
}

func oidWithin(oid string, from string, includeFrom bool, to string) bool {
	oidBytes, fromBytes, toBytes := []byte(oid), []byte(from), []byte(to)

	fromCompare := bytes.Compare(fromBytes, oidBytes)
	toCompare := bytes.Compare(toBytes, oidBytes)

	return (fromCompare == -1 || (fromCompare == 0 && includeFrom)) && (toCompare == 1)
}
