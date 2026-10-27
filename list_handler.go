// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import (
	"context"
	"sort"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

// ListHandler is a helper that takes a list of oids and implements
// a default behaviour for that list.
type ListHandler struct {
	oids  []value.OID
	items map[string]*ListItem
}

// Add adds a list item for the provided oid and returns it.
func (l *ListHandler) Add(oid string) *ListItem {
	if l.items == nil {
		l.items = make(map[string]*ListItem)
	}

	parsedOID := value.MustParseOID(oid)
	l.oids = append(l.oids, parsedOID)
	value.SortOIDs(l.oids)
	item := &ListItem{}
	l.items[oid] = item
	return item
}

// AddUnsorted adds a list item like Add but does not keep the OID list
// sorted, so building a large list is not quadratic. Call Sort once after
// the last AddUnsorted and before the handler serves GetNext.
func (l *ListHandler) AddUnsorted(oid string) *ListItem {
	if l.items == nil {
		l.items = make(map[string]*ListItem)
	}

	l.oids = append(l.oids, value.MustParseOID(oid))
	item := &ListItem{}
	l.items[oid] = item
	return item
}

// Sort sorts the OID list after a series of AddUnsorted calls.
func (l *ListHandler) Sort() {
	value.SortOIDs(l.oids)
}

// Get tries to find the provided oid and returns the corresponding value.
func (l *ListHandler) Get(ctx context.Context, oid value.OID) (value.OID, pdu.VariableType, any, error) {
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
func (l *ListHandler) GetNext(ctx context.Context, from value.OID, includeFrom bool, to value.OID) (value.OID, pdu.VariableType, any, error) {
	if l.items == nil {
		return nil, pdu.VariableTypeNoSuchObject, nil, nil
	}

	// The OIDs are sorted, so the first one past from is the only
	// candidate: every later OID is larger and cannot be below to either.
	i := sort.Search(len(l.oids), func(i int) bool {
		c := value.CompareOIDs(l.oids[i], from)
		return c == 1 || (c == 0 && includeFrom)
	})
	if i < len(l.oids) && oidWithin(l.oids[i], from, includeFrom, to) {
		return l.Get(ctx, l.oids[i])
	}

	return nil, pdu.VariableTypeNoSuchObject, nil, nil
}

func oidWithin(oid value.OID, from value.OID, includeFrom bool, to value.OID) bool {
	fromCompare := value.CompareOIDs(from, oid)
	toCompare := value.CompareOIDs(to, oid)

	return (fromCompare == -1 || (fromCompare == 0 && includeFrom)) && (toCompare == 1)
}
