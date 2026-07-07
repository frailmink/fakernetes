package cli

import (
	"github.com/spf13/cobra"

	"github.com/frailmink/fakernetes/internal/fakelet"
)

func NewFakeletSubCommand(rootCmd *cobra.Command, rootCfg *RootCfg) {
 	fakeletCmd := &cobra.Command{
		Use:   "fakelet [options]",
		Short: "fakelet will monitor the pods and start up the containers",
		RunE: commandExecWrapper(rootCfg, func(cmd *cobra.Command, args []string) error {
			return fakelet.Execute(rootCfg.Logger)
		}),
	}

	rootCmd.AddCommand(fakeletCmd)
}
