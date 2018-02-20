// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package agentx_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/posteo/go-agentx"
	"github.com/posteo/go-agentx/value"
	"gopkg.in/errgo.v1"
)

type environment struct {
	client *agentx.Client
}

var (
	e *environment

	baseOID = value.MustParseOID("1.3.6.1.4.1.45995")
)

func TestMain(m *testing.M) {
	e = &environment{}
	e.client = &agentx.Client{
		Net:     "tcp",
		Address: "localhost:705",
		Timeout: 60 * time.Second,
		NameOID: value.MustParseOID("1.3.6.1.4.1.45995"),
		Name:    "test client",
	}

	if err := e.client.Open(); err != nil {
		log.Fatalf(errgo.Details(err))
	}

	result := m.Run()

	if err := e.client.Close(); err != nil {
		log.Fatalf(errgo.Details(err))
	}

	os.Exit(result)
}
