/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"needs-a-name/pkg/versioning/semver"
)

var bumpCmd = &cobra.Command{
	Use:   "bump VERSION",
	Short: "Bumps the specified version number",
	Long: `Bumps the specified version number with the specified version level.
Can also be used to bump multiple levels with the --number flag`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		version, err := semver.Bump(args[0], versionLevel, numberOfVersions)
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}
		fmt.Println(version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(bumpCmd)
}
