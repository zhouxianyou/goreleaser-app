package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Populated by goreleaser during build
	version = "master"
	commit  = "none"
	date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("goreleaser-app has version %s built from %s on %s\n", version, commit, date)
	},
}
