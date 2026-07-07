package cli

import (
	"github.com/spf13/cobra"
)

var fakeletCmd = &cobra.Command{
	Use:   "fakelet [options]",
	Short: "fakelet will monitor the pods and start up the containers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init()  {
	rootCmd.AddCommand(fakeletCmd)
}
