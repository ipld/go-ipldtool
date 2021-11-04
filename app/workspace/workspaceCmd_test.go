package workspace_test

import (
	"testing"

	"github.com/ipld/go-ipldtool/app/testutil"
)

func TestRead(t *testing.T) {
	testutil.TestExecSpec(t, "../../docs/workspace.md")
}
