// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx

import (
	"context"
	"slices"
	"sort"

	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

// BulkHandler serves a table that is built in one go and then only read,
// such as a snapshot of tens of thousands of rows that is rebuilt and
// swapped as a whole. Unlike ListHandler it keeps no string keys: Add
// appends a numeric OID, Sort orders the rows once, and Get and GetNext
// binary-search. Build it completely, call Sort, then hand it to a session;
// it must not be modified while it is being served.
type BulkHandler struct {
	rows []bulkRow
}

type bulkRow struct {
	oid value.OID
	ListItem
}

// Grow makes room for n more rows.
func (b *BulkHandler) Grow(n int) {
	b.rows = slices.Grow(b.rows, n)
}

// Add appends a row. The handler keeps the OID slice; do not modify it.
func (b *BulkHandler) Add(oid value.OID, t pdu.VariableType, v any) {
	b.rows = append(b.rows, bulkRow{oid: oid, ListItem: ListItem{Type: t, Value: v}})
}

// Len returns the number of rows.
func (b *BulkHandler) Len() int {
	return len(b.rows)
}

// Sort orders the rows by OID. Of several rows with the same OID the last
// one added is kept, as ListHandler.Get would return.
func (b *BulkHandler) Sort() {
	slices.SortStableFunc(b.rows, func(x, y bulkRow) int { return value.CompareOIDs(x.oid, y.oid) })
	out := b.rows[:0]
	for i, r := range b.rows {
		if i+1 < len(b.rows) && value.CompareOIDs(r.oid, b.rows[i+1].oid) == 0 {
			continue
		}
		out = append(out, r)
	}
	b.rows = out
}

// Get returns the row with exactly this OID.
func (b *BulkHandler) Get(ctx context.Context, oid value.OID) (value.OID, pdu.VariableType, any, error) {
	i := sort.Search(len(b.rows), func(i int) bool { return value.CompareOIDs(b.rows[i].oid, oid) >= 0 })
	if i < len(b.rows) && value.CompareOIDs(b.rows[i].oid, oid) == 0 {
		return oid, b.rows[i].Type, b.rows[i].Value, nil
	}
	return nil, pdu.VariableTypeNoSuchObject, nil, nil
}

// GetNext returns the first row after from (or at from, with includeFrom)
// that lies below to, with the same semantics as ListHandler.GetNext.
func (b *BulkHandler) GetNext(ctx context.Context, from value.OID, includeFrom bool, to value.OID) (value.OID, pdu.VariableType, any, error) {
	i := sort.Search(len(b.rows), func(i int) bool {
		c := value.CompareOIDs(b.rows[i].oid, from)
		return c == 1 || (c == 0 && includeFrom)
	})
	if i < len(b.rows) && oidWithin(b.rows[i].oid, from, includeFrom, to) {
		return b.rows[i].oid, b.rows[i].Type, b.rows[i].Value, nil
	}
	return nil, pdu.VariableTypeNoSuchObject, nil, nil
}
