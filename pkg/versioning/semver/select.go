/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package semver

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"go.uber.org/zap"
	"slices"
)

func Select(tags []string, numberOfVersions int, versionLevel string) (selectedVersions []*semver.Version) {
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
			selectedVersions = append(selectedVersions, version)
			if len(selectedVersions) == numberOfVersions {
				break
			}
		}
	}
	slices.Reverse(selectedVersions)
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
