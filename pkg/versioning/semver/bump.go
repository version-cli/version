/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package semver

import (
	"strings"

	"github.com/coreos/go-semver/semver"
	"go.uber.org/zap"
)

func Bump(versionNumber string, versionLevel string, count int) (version string, err error) {
	parsedVersion, err := parse(versionNumber)
	if err != nil {
		zap.L().Error("Error while parsing version", zap.String("versionNumber", versionNumber))
		return "", err
	}
	for i := 0; i < count; i++ {
		parsedVersion = bumpVersion(versionLevel, parsedVersion)
	}
	version = parsedVersion.String()
	if strings.HasPrefix(versionNumber, "v") {
		version = "v" + version
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
