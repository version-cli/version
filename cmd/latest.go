/*
Copyright © 2024 Koen van Zuijlen <8818390+kvanzuijlen@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"github.com/kvanzuijlen/version/pkg/datasource"
	"github.com/kvanzuijlen/version/pkg/versioning/semver"

	"github.com/spf13/cobra"
)

var (
	datasourceName string
	latestCmd      = &cobra.Command{
		Use:   "latest DEPNAME",
		Short: "Retrieves the latest (n) versions of a dependency from a datasource",
		Long: `Retrieves the specified number of latest versions of a dependency from the specified datasource.
If a different version level is specified, this command will return the latest number of versions of that version level.`,
		Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ds, err := datasource.Get(datasourceName)
			if err != nil {
				return err
			}
			tags, err := ds.Latest(args[0])
			if err != nil {
				return err
			}
			versions := semver.Select(tags, numberOfVersions, versionLevel)
			for _, version := range versions {
				fmt.Println(version)
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(latestCmd)
	latestCmd.PersistentFlags().StringVarP(&datasourceName, "datasource", "d", "", "datasource to use")
	err := latestCmd.MarkFlagRequired("datasource")
	if err != nil {
		return
	}
}
