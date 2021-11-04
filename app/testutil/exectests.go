package testutil

import (
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/frankban/quicktest"
	"github.com/ipld/go-ipldtool/app"
	"github.com/warpfork/go-testmark"
	"github.com/warpfork/go-testmark/testexec"
)

func TestExecSpec(t *testing.T, specFile string) {
	t.Helper()

	// Just real quick compile the whole app.  We need this so we can test it in scripts.
	//  Hardcoded output path seems to work in practice.  (You'd think it'd be racey, but apparently it's survivable.)
	//  Hardcoded compile path also works, for now, by coincidence, because all users are at the same relative package depth.
	os.MkdirAll("/tmp/ipld-test/bin/", 0755)
	if err := exec.Command("go", "build", "-o", "/tmp/ipld-test/bin/ipld", "../../cmd/ipld/ipld.go").Run(); err != nil {
		t.Fatalf("failed to build the command: %s", err)
	}

	// Read the spec file.
	doc, err := testmark.ReadFile(specFile)
	if err != nil {
		t.Fatalf("spec file parse failed?!: %s", err)
	}

	// Make the testexec config structure.
	//  If using ExecFn, it'll use the app's main method (i.e. not spawn actual subprocesses).
	//  If using ScriptFn, we'll do some PATH env hijinx to get the command in place first.
	//  Our assertions use quicktest.CmpEqual for diffing.
	//  And of course we want to accumulate any patches, if we're being run in regen mode.
	tcfg := testexec.Tester{
		ExecFn: app.Main,
		ScriptFn: func(script string, stdin io.Reader, stdout, stderr io.Writer) (exitcode int, oshit error) {
			return testexec.ScriptFn_ExecBash("export PATH=$PATH:/tmp/ipld-test/bin/; export IPLDTOOL_NOHOME=true;\n"+script, stdin, stdout, stderr)
		},
		AssertFn: func(t *testing.T, actual, expect string) {
			quicktest.Assert(t, actual, quicktest.CmpEquals(), expect)
		},
		Patches: &testmark.PatchAccumulator{},
	}

	// Data hunk in this spec file are in "directories" of a test scenario each.
	doc.BuildDirIndex()
	for _, dir := range doc.DirEnt.ChildrenList {
		t.Run(dir.Name, func(t *testing.T) {
			tcfg.Test(t, dir)
		})
	}

	// Write back any patches.
	tcfg.Patches.WriteFileWithPatches(doc, specFile)
}
