package schema_test

import (
	"runtime"
	"testing"

	"github.com/ipld/go-ipldtool/app/testutil"
)

func TestSchema(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip()
	}
	testutil.TestExecSpec(t, "../../docs/schema.md")
}
