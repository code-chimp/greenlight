package vcs

import (
	"fmt"
	"runtime/debug"
)

// Revision returns the current VCS revision of the build.
// It reads the build information to extract the VCS revision and checks if the source code was modified.
//
// Returns:
//
//	string - The VCS revision. If the source code was modified, "-dirty" is appended to the revision.
func Revision() string {
	var revision string
	var modified bool

	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range info.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value
			case "vcs.modified":
				modified = s.Value == "true"
			}
		}
	}

	if modified {
		return fmt.Sprintf("%s-dirty", revision)
	}

	return revision
}
