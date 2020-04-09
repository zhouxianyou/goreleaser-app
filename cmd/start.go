package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/zhouxianyou/goreleaser-app/packages"
)

// startCmd is starting node
var startCmd = &cobra.Command{
	Use:   "packages",
	Short: "Starting node",
	Run: func(cmd *cobra.Command, args []string) {
		packages.Start()
	},
}

func init() {
	time.Local = time.UTC
}
