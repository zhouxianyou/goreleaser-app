
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goreleaser",
	Short: "application",
}

func init() {
	rootCmd.AddCommand(
		versionCmd,
		startCmd,
	)

}

// Execute executes rootCmd command.
// This is called by main.main(). It only needs to happen once to the rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Executing root command")
	}
}
