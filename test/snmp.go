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
