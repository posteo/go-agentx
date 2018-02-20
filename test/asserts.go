// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package test

import (
	"reflect"
	"testing"

	"gopkg.in/errgo.v1"
)

func AssertNoError(tb testing.TB, err error) {
	if err != nil {
		tb.Fatalf(errgo.Details(err))
	}
}

func AssertEquals(tb testing.TB, expected, actucal interface{}) {
	if !reflect.DeepEqual(expected, actucal) {
		tb.Fatalf("expected %#v, got %#v", expected, actucal)
	}
}

func AssertNotEquals(tb testing.TB, expected, actucal interface{}) {
	if reflect.DeepEqual(expected, actucal) {
		tb.Fatalf("expected not %#v, got %#v", expected, actucal)
	}
}
