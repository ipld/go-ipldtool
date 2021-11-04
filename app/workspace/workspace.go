package workspace

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	ipldtoolerr "github.com/ipld/go-ipldtool/errors"
)

const MagicWorkspaceDirname = ".ipld"

// Find the nearest workspace, based on search around cwd and env vars.
//
// The string returned is a path to a directory that contains an '.ipld' directory.
// (E.g. the string returned does not include the '.ipld' segment itself.)
//
// The search operations, in order, are:
//
//   - accept an env var overriding it: `$IPLDTOOL_WORKSPACE`.
//   - look at the present dir so see if it contains an `.ipld` dir.
//   - look at parent dirs, recursively, to see if it contains an `.ipld` dir.
//   - if nothing is found yet: if `$IPLDTOOL_NOHOME` is set, stop, none found.
//     (Necessary to have an option for this so we can handle readonly fs without a fuss.)
//   - fallback to `$HOME/.ipld` as a workspace (even if we have to create it).
//
// Edge cases: if the '.ipld' path exists, but is a file, we don't comment on it.
// (You should get errors from subsequent operations that are informative enough anyway.)
// If we fell back to an '.ipld' dir in the user homedir, this function may create it.
//
// Errors:
//
//  - ipldtool-workspace-not-found -- if we tried everything and can't find a workspace.
//  - ipldtool-error-no-cwd -- if the cwd can't be found!
//  - ipldtool-error-io -- if there's an io error during the search (permission denied, etc).
//
func Find() (string, error) {
	// If the override var is present: that's it.
	if override := os.Getenv("IPLDTOOL_WORKSPACE"); override != "" {
		return override, nil
	}

	// Search up, starting from cwd.
	cwd, err := os.Getwd()
	if err != nil {
		return "", ipldtoolerr.Newf("ipldtool-error-no-cwd", "%s", err)
	}
	searchAt := cwd
	for {
		// Probe fairly blindly.
		//  (We're mostly expecting a directory, but don't actually check that here.)
		f, err := os.Open(filepath.Join(searchAt, MagicWorkspaceDirname))
		if f != nil {
			f.Close()
		}
		if err == nil { // no error?  Found it!
			return searchAt, nil
		}
		if errors.Is(err, fs.ErrNotExist) { // no such thing?  oh well.  pop a segment and keep looking.
			searchAt = filepath.Dir(searchAt)
			// If popping a searchAt segment got us down to nothing,
			//  and we didn't find anything here either,
			//   that's it: this part of the search is done.  Break out of the loop.
			if searchAt == "/" || searchAt == "." {
				break
			}
			// ... otherwise: continue, with popped searchAt.
			continue
		}
		// You're still here?  That means there's an error, but of some unpleasant kind.
		//  Whatever this error is, our search has blind spots: error out.
		return "", ipldtoolerr.Newf("ipldtool-error-io", "error during search for workspace: %s", err)
	}

	// Still nada?  Check the homedir.  Unless we were told not to, of course.
	// And make it, if it doesn't exist.
	if nohome := os.Getenv("IPLDTOOL_NOHOME"); nohome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", ipldtoolerr.Newf("ipldtool-error-io", "error during search for workspace: can't find homedir: %s", err)
		}
		if err := os.Mkdir(filepath.Join(home, MagicWorkspaceDirname), 0755); err != nil {
			if errors.Is(err, fs.ErrExist) {
				return home, nil
			}
			return "", ipldtoolerr.Newf("ipldtool-error-io", "could not make workspace dir in homedir: %s", err)
		}
		return home, nil
	}

	// All options exhausted.  Report a not found.
	return "", ipldtoolerr.Newf("ipldtool-workspace-not-found", "no workspace marker (an '.ipld' dir) found while searching up from %q", cwd)
}
