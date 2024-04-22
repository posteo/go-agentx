// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Yamu-OSS/go-agentx/value"
)

func TestSession(t *testing.T) {
	e := setUpTestEnvironment(t)
	defer e.tearDown()

	t.Run("Open", func(t *testing.T) {
		session, err := e.client.Session()
		require.NoError(t, err)
		defer session.Close()

		assert.NotEqual(t, 0, session.ID())
	})

	t.Run("Close", func(t *testing.T) {
		session, err := e.client.Session()
		require.NoError(t, err)

		require.NoError(t, session.Close())
	})

	t.Run("Register", func(t *testing.T) {
		session, err := e.client.Session()
		require.NoError(t, err)
		defer session.Close()

		baseOID := value.MustParseOID("1.3.6.1.4.1.45995")

		require.NoError(t,
			session.Register(127, baseOID))

		require.NoError(t,
			session.Unregister(127, baseOID))
	})
}
