// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx_test

import (
	"io"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Yamu-OSS/go-agentx"
	"github.com/Yamu-OSS/go-agentx/value"
)

type environment struct {
	client   *agentx.Client
	tearDown func()
}

func setUpTestEnvironment(tb testing.TB) *environment {
	cmd := exec.Command("snmpd", "-Lo", "-f", "-c", "snmpd.conf")

	stdout, err := cmd.StdoutPipe()
	require.NoError(tb, err)
	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	log.Printf("run: %s", cmd)
	require.NoError(tb, cmd.Start())
	time.Sleep(500 * time.Millisecond)

	client, err := agentx.Dial("tcp", "127.0.0.1:30705")
	require.NoError(tb, err)
	client.Timeout = 60 * time.Second
	client.NameOID = value.MustParseOID("1.3.6.1.4.1.45995")
	client.Name = "test client"

	return &environment{
		client: client,
		tearDown: func() {
			require.NoError(tb, client.Close())
			require.NoError(tb, cmd.Process.Kill())
		},
	}
}
