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
