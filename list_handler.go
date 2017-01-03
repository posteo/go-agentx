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

package agentx

import (
	"sort"

	"github.com/martinclaro/go-oidsort"
	"github.com/martinclaro/go-agentx/pdu"
	"github.com/martinclaro/go-agentx/value"
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
	sort.Sort(oidsort.ByOidString(l.oids))
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
	fromCompare := oidsort.CompareOIDs(from, oid)
	toCompare := oidsort.CompareOIDs(to, oid)

	return (fromCompare == -1 || (fromCompare == 0 && includeFrom)) && (toCompare == 1)
}
