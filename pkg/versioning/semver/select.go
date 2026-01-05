/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package semver

import (
	"fmt"
	"slices"
	"strings"

	"github.com/coreos/go-semver/semver"
	"go.uber.org/zap"
)

func Select(tags []string, count int, versionLevel string) (selectedVersions []string) {
	versions := semverSort(tags)

	versionMap := make(map[string]*semver.Version)
	for _, version := range versions {
		key, err := getVersionKeyForLevel(version, versionLevel)
		if err != nil {
			zap.L().Debug(err.Error())
			continue
		}
		if _, exists := versionMap[key]; !exists {
			versionMap[key] = version
			selectedVersions = append(selectedVersions, version.String())
			if len(selectedVersions) == count {
				break
			}
		}
	}
	slices.Reverse(selectedVersions)
	// Check if any of the tags had a 'v' prefix, if so, we should probably return versions with 'v' prefix
	// This is a heuristic, but it's probably good enough
	hasVPrefix := false
	for _, tag := range tags {
		if strings.HasPrefix(tag, "v") {
			hasVPrefix = true
			break
		}
	}

	if hasVPrefix {
		for i, version := range selectedVersions {
			selectedVersions[i] = "v" + version
		}
	}

	return selectedVersions
}

func getVersionKeyForLevel(version *semver.Version, versionLevel string) (string, error) {
	switch versionLevel {
	case "major":
		return fmt.Sprintf("%d", version.Major), nil
	case "minor":
		return fmt.Sprintf("%d.%d", version.Major, version.Minor), nil
	case "patch":
		return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch), nil
	default:
		return "", fmt.Errorf("invalid semantic version: %s", version)
	}
}
