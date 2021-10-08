package schema_test

import (
	"testing"

	"github.com/ipld/go-ipldtool/app/testutil"
)

func TestSchema(t *testing.T) {
	testutil.TestExecSpec(t, "../../docs/schema.md")
}
