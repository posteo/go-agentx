// Copyright 2018 The agentx authors
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
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
