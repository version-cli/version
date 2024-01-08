/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package semver

import "github.com/coreos/go-semver/semver"

func parse(versionNumber string) (version *semver.Version, err error) {
	return semver.NewVersion(versionNumber)
}
