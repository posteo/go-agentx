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
