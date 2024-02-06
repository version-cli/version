/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package semver

import (
	"github.com/coreos/go-semver/semver"
	"go.uber.org/zap"
)

func Bump(versionNumber string, versionLevel string, count int) (version *semver.Version, err error) {
	version, err = parse(versionNumber)
	if err != nil {
		zap.L().Error("Error while parsing version", zap.String("versionNumber", versionNumber))
		return nil, err
	}
	for i := 0; i < count; i++ {
		version = bumpVersion(versionLevel, version)
	}
	return version, nil
}

func bumpVersion(versionLevel string, version *semver.Version) *semver.Version {
	switch versionLevel {
	case "major":
		version.BumpMajor()
	case "minor":
		version.BumpMinor()
	case "patch":
		version.BumpPatch()
	}
	return version
}
