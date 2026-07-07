package cli

import (
	"log/slog"
	"github.com/spf13/cobra"
)

func commandExecWrapper(rootCfg *RootCfg, f func(*cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if rootCfg.debug {
			rootCfg.LogLevelVar.Set(slog.LevelDebug)
		}
		return f(cmd, args)
	}
}
