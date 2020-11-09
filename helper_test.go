package agentx_test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func SNMPGet(tb testing.TB, oid string) string {
	cmd := exec.Command("snmpget", "-v2c", "-cpublic", "-On", "127.0.0.1:30161", oid)
	output, err := cmd.CombinedOutput()
	require.NoError(tb, err)
	return strings.TrimSpace(string(output))
}

func SNMPGetNext(tb testing.TB, oid string) string {
	cmd := exec.Command("snmpgetnext", "-v2c", "-cpublic", "-On", "127.0.0.1:30161", oid)
	output, err := cmd.CombinedOutput()
	require.NoError(tb, err)
	return strings.TrimSpace(string(output))
}

func SNMPGetBulk(tb testing.TB, oid string, nonRepeaters, maxRepetitions int) string {
	cmd := exec.Command("snmpbulkget", "-v2c", "-cpublic", "-On", fmt.Sprintf("-Cn%d", nonRepeaters), fmt.Sprintf("-Cr%d", maxRepetitions), "127.0.0.1:30161", oid)
	output, err := cmd.CombinedOutput()
	require.NoError(tb, err)
	return strings.TrimSpace(string(output))
}
