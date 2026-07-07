package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/frailmink/fakernetes/internal/cli"
)

func main() {
	logLevel := new(slog.LevelVar)
	logLevel.Set(slog.LevelInfo)
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	}))

	rootCfg := &cli.RootCfg{
		LogLevelVar: logLevel,
		Logger:      logger,
	}

	rootCmd := cli.NewRootCmd(rootCfg)

	if err := cli.Execute(rootCmd, rootCfg); err != nil {
		fmt.Printf("An error occurred when executing the command: %v", err)
		os.Exit(1)
	}
}
