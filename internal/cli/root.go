package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fakernetes <sub-command> [options]",
	Short: "fakernetes is a cli which can be used to run a cluster",
	Long:  "fakernetes is a cli which can be used to run a cluster - It will allow you to run all the components necessary for the cluster to be operational",
}

func Execute() error {
	return rootCmd.Execute()
}
