/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package cmd

import (
	"os"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var (
	numberOfVersions    int
	major, minor, patch bool
	versionLevel        string
	useSemver           bool
	VERSION             string
	rootCmd             = &cobra.Command{
		Use:              "version",
		Short:            "A tool to handle version numbers",
		Long:             `A tool to handle version numbers, i.e. bumping versions and retrieving the latest version of a dependency.`,
		PersistentPreRun: setVersionLevel,
		Version:          VERSION,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&numberOfVersions, "count", "c", 1, "Number of versions")
	rootCmd.PersistentFlags().BoolVarP(&major, "major", "M", false, "Set version level to major")
	rootCmd.PersistentFlags().BoolVarP(&minor, "minor", "m", false, "Set version level to minor")
	rootCmd.PersistentFlags().BoolVarP(&patch, "patch", "p", true, "Set version level to patch")
	rootCmd.MarkFlagsMutuallyExclusive("major", "minor", "patch")
	rootCmd.PersistentFlags().BoolVarP(&useSemver, "semver", "S", true, "Use semver as the versioning type")
}

func setVersionLevel(_ *cobra.Command, _ []string) {
	switch {
	case major:
		versionLevel = "major"
	case minor:
		versionLevel = "minor"
	case patch:
		versionLevel = "patch"
	default:
		versionLevel = "patch"
	}

	zap.L().Debug("Starting...",
		zap.String("version level", versionLevel),
		zap.Int("number of versions", numberOfVersions),
	)
}
