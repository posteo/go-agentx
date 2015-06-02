package agentx_test

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/errgo.v1"
)

func SNMPGet(tb testing.TB, oid string) string {
	command := exec.Command("/usr/bin/snmpget", "-v2c", "-cpublic", "-On", "localhost", oid)
	output, err := command.CombinedOutput()
	AssertNoError(tb, err)
	return strings.TrimSpace(string(output))
}

func SNMPGetNext(tb testing.TB, oid string) string {
	command := exec.Command("/usr/bin/snmpgetnext", "-v2c", "-cpublic", "-On", "localhost", oid)
	output, err := command.CombinedOutput()
	err = nil
	AssertNoError(tb, err)
	return strings.TrimSpace(string(output))
}

func SNMPGetBulk(tb testing.TB, oid string, nonRepeaters, maxRepetitions int) string {
	command := exec.Command("/usr/bin/snmpbulkget", "-v2c", "-cpublic", "-On", fmt.Sprintf("-Cn%d", nonRepeaters), fmt.Sprintf("-Cr%d", maxRepetitions), "localhost", oid)
	output, err := command.CombinedOutput()
	err = nil
	AssertNoError(tb, err)
	return strings.TrimSpace(string(output))
}

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
