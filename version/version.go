package version

import (
	"strings"
)

var (
	// The git commit that was compiled. These will be filled in by the
	// compiler.
	GitCommit string

	// The main version number that is being run at the moment.
	//
	// Version must conform to the format expected by github.com/hashicorp/go-version
	// for tests to work.
	Version = "1.12.0"

	// A pre-release marker for the version. If this is "" (empty string)
	// then it means that it is a final release. Otherwise, this is a pre-release
	// such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = "fly3"
)

// GetHumanVersion composes the parts of the version in a way that's suitable
// for displaying to humans.
func GetHumanVersion() string {
	version := Version
	release := VersionPrerelease

	if release != "" {
		suffix := "-" + release
		if !strings.HasSuffix(version, suffix) {
			// if we tagged a prerelease version then the release is in the version already
			version += suffix
		}
	}

	// Strip off any single quotes added by the git information.
	return strings.ReplaceAll(version, "'", "")
}
