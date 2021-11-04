package workspace

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"

	ipldtoolerr "github.com/ipld/go-ipldtool/errors"
)

var Cmd_Workspace = &cli.Command{
	Name:     "workspace",
	Category: "basic",
	Usage:    "Create, configure, or interogate a workspace for the ipldtool.  (You'll need a workspace for any of the stateful commands.)",
	Subcommands: []*cli.Command{{
		Name:  "new",
		Usage: "Creates the local filesystem markers for a new workspace.",
		Flags: []cli.Flag{
			// Some, eventually, probably -- but for now not even the storage engine is configurable, yet.
		},
		Action: Action_WorkspaceNew,
	}, {
		Name:   "find",
		Usage:  "Tells you what the current workspace is.",
		Action: Action_WorkspaceFind,
	}},
}

// Action_WorkspaceNew is the 'ipld workspace new' command.
//
// Errors:
//
//   - ipldtool-error-invalid-args -- for incomprehensible or invalid arguments.
//   - ipldtool-error-io -- if there's an io error (permission denied, readonly disk, etc).
//
// It is not an error if the workspace already exists.
func Action_WorkspaceNew(args *cli.Context) error {
	// Parse positional args.
	var targetDir string
	switch args.Args().Len() {
	case 0:
		targetDir = ""
	case 1:
		targetDir = args.Args().Get(0)
	default:
		return ipldtoolerr.Newf(ipldtoolerr.ErrCode_InvalidArgs, "'workspace new' command needs zero or one positional argument")
	}

	// Make the directory exist.
	workspaceDir := filepath.Join(targetDir, MagicWorkspaceDirname)
	if err := os.MkdirAll(workspaceDir, 0755); err != nil {
		return ipldtoolerr.Newf("ipldtool-error-io", "could not create new workspace: %s", err)
	}

	// That's it!  There's no other required configuration.
	return nil
}

// Action_WorkspaceFind is the 'ipld workspace find' command.
// It prints out the path to the workspace,
// or nothing if one isn't found.
//
// Errors:
//
//   - ipldtool-error-invalid-args -- for incomprehensible or invalid arguments.
//   - ipldtool-error-no-cwd -- if the cwd can't be found!
//   - ipldtool-error-io -- if there's an io error (permission denied, readonly disk, etc).
//   - ipldtool-workspace-not-found -- FIXME: this isn't actually possible.
//
// It is not an error if the workspace already exists.
func Action_WorkspaceFind(args *cli.Context) error {
	// Parse positional args.
	switch args.Args().Len() {
	case 0:
		wspath, err := Find()
		switch err.(*ipldtoolerr.Error).Code() {
		case "":
			return nil
		case "ipldtool-workspace-not-found":
			return nil // Fine.  Silence is our answer, then, in this command.  No problem.
		default:
			return err // (Wish: the error analysis tool could subtract the cases above from codes still possible in `err` here.)
		}
		fmt.Fprintf(args.App.Writer, "%s\n", wspath)
		return nil
	default:
		// I flirted with the idea of accepting a positional arg for where to start the search from,
		// but this seems like a bad idea; or if we do it, that should be a consistent arg for all commands globally.
		// For now, your options are changing CWD, or overriding the search entirely via IPLDTOOL_WORKSPACE.
		return ipldtoolerr.Newf(ipldtoolerr.ErrCode_InvalidArgs, "'workspace find' command needs zero or one positional argument")
	}
}
