package agentx_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/juju/errgo"
	"github.com/posteo/go-agentx"
)

type environment struct {
	client *agentx.Client
}

var e *environment

func TestMain(m *testing.M) {
	e = &environment{}
	e.client = &agentx.Client{
		Net:     "tcp",
		Address: "localhost:705",
		Timeout: 60 * time.Second,
		NameOID: "1.3.6.1.4.1.8072",
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
