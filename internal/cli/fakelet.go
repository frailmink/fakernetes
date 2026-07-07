package cli

import (
	"github.com/spf13/cobra"

	"github.com/frailmink/fakernetes/internal/fakelet"
)

var fakeletCmd = &cobra.Command{
	Use:   "fakelet [options]",
	Short: "fakelet will monitor the pods and start up the containers",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fakelet.Execute()
	},
}

func init()  {
	rootCmd.AddCommand(fakeletCmd)
}
