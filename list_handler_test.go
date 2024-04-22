// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Yamu-OSS/go-agentx"
	"github.com/Yamu-OSS/go-agentx/pdu"
	"github.com/Yamu-OSS/go-agentx/value"
)

func TestListHandler(t *testing.T) {
	e := setUpTestEnvironment(t)
	defer e.tearDown()

	session, err := e.client.Session()
	require.NoError(t, err)
	defer session.Close()

	lh := &agentx.ListHandler{}
	i := lh.Add("1.3.6.1.4.1.45995.3.1")
	i.Type = pdu.VariableTypeOctetString
	i.Value = "test"

	i = lh.Add("1.3.6.1.4.1.45995.3.3")
	i.Type = pdu.VariableTypeOctetString
	i.Value = "test2"

	i = lh.Add("1.3.6.1.4.1.45995.3.4")
	i.Type = pdu.VariableTypeOctetString
	i.Value = "test3"

	session.Handler = lh

	baseOID := value.MustParseOID("1.3.6.1.4.1.45995")

	require.NoError(t, session.Register(127, baseOID))
	defer session.Unregister(127, baseOID)

	t.Run("Get", func(t *testing.T) {
		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
			SNMPGetBatch(t, "1.3.6.1.4.1.45995.3.1"))

		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.2 = No Such Object available on this agent at this OID",
			SNMPGetBatch(t, "1.3.6.1.4.1.45995.3.2"))

		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.3 = STRING: \"test2\"",
			SNMPGetBatch(t, "1.3.6.1.4.1.45995.3.3"))

		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"\n.1.3.6.1.4.1.45995.3.3 = STRING: \"test2\"",
			SNMPGetBatch(t, "1.3.6.1.4.1.45995.3.1", "1.3.6.1.4.1.45995.3.3"))

		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"\n.1.3.6.1.4.1.45995.3.2 = No Such Object available on this agent at this OID",
			SNMPGetBatch(t, "1.3.6.1.4.1.45995.3.1", "1.3.6.1.4.1.45995.3.2"))

		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.2 = No Such Object available on this agent at this OID\n.1.3.6.1.4.1.45995.3.3 = STRING: \"test2\"",
			SNMPGetBatch(t, "1.3.6.1.4.1.45995.3.2", "1.3.6.1.4.1.45995.3.3"))

		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"\n.1.3.6.1.4.1.45995.3.3 = STRING: \"test2\"\n.1.3.6.1.4.1.45995.3.4 = STRING: \"test3\"",
			SNMPGetBatch(t, "1.3.6.1.4.1.45995.3.1", "1.3.6.1.4.1.45995.3.3", "1.3.6.1.4.1.45995.3.4"))
	})

	t.Run("GetNext", func(t *testing.T) {
		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
			SNMPGetNext(t, "1.3.6.1.4.1.45995.3.0"))

		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
			SNMPGetNext(t, "1.3.6.1.4.1.45995.3"))

	})

	t.Run("GetBulk", func(t *testing.T) {
		assert.Equal(t,
			".1.3.6.1.4.1.45995.3.1 = STRING: \"test\"",
			SNMPGetBulk(t, "1.3.6.1.4.1.45995.3.0", 0, 1))
	})
}
