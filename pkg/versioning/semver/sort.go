/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package semver

import (
	"sort"

	"github.com/coreos/go-semver/semver"
)

func semverSort(list []string) (versions semver.Versions) {
	for _, version := range list {
		semVersion, err := parse(version)
		if err != nil {
			continue
		}
		versions = append(versions, semVersion)
	}
	sort.Sort(sort.Reverse(versions))
	return versions
}
