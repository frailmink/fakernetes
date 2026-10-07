package cli

import (
	"log/slog"

	"github.com/spf13/cobra"
)

type RootCfg struct {
	Logger      *slog.Logger
	LogLevelVar *slog.LevelVar
	debug       bool
}

func NewRootCmd(cfg *RootCfg) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "fakernetes <sub-command> [options]",
		Short: "fakernetes is a cli which can be used to run a cluster",
		Long:  "fakernetes is a cli which can be used to run a cluster - It will allow you to run all the components necessary for the cluster to be operational",
	}

	rootCmd.PersistentFlags().BoolVarP(&cfg.debug, "debug", "d", false, "Print debug logs")

	return rootCmd
}

func Execute(rootCmd *cobra.Command, rootCfg *RootCfg) error {
	newFakeletSubCommand(rootCmd, rootCfg)

	return rootCmd.Execute()
}
