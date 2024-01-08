/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package semver

import "github.com/coreos/go-semver/semver"

func Bump(versionNumber string, versionLevel string, numberOfVersions int) (version *semver.Version, err error) {
	version, err = parse(versionNumber)
	for i := 0; i < numberOfVersions; i++ {
		version = bumpVersion(versionLevel, version)
	}
	return version, nil
}

func bumpVersion(versionLevel string, version *semver.Version) *semver.Version {
	switch versionLevel {
	case "major":
		version.BumpMajor()
		break
	case "minor":
		version.BumpMinor()
		break
	case "patch":
		version.BumpPatch()
		break
	}
	return version
}
