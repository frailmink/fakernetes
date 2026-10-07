package cli

import (
	"github.com/spf13/cobra"

	"github.com/frailmink/fakernetes/internal/fakelet"
)

func newFakeletSubCommand(rootCmd *cobra.Command, rootCfg *RootCfg) {
	fakeletConfig := fakelet.FakeletConfig{
		Logger: rootCfg.Logger,
		ContainerdSocket: "/run/containerd/containerd.sock",
	}

 	fakeletCmd := &cobra.Command{
		Use:   "fakelet [options]",
		Short: "fakelet will monitor the pods and start up the containers",
		RunE: commandExecWrapper(rootCfg, func(cmd *cobra.Command, args []string) error {
			// This function is a closure, the fakeletConfig var's pointer is conceptually captured and when this function is called
			// the value of fakeletConfig reads the newest changes
			return fakelet.Execute(fakeletConfig)
		}),
	}

	fakeletCmd.Flags().StringVarP(&fakeletConfig.ContainerdSocket, "soc", "s", fakeletConfig.ContainerdSocket, "Directory pointing to the containerd socket")

	rootCmd.AddCommand(fakeletCmd)
}
