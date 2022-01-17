package workspace_test

import (
	"runtime"
	"testing"

	"github.com/ipld/go-ipldtool/app/testutil"
)

func TestRead(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip()
	}
	testutil.TestExecSpec(t, "../../docs/workspace.md")
}
