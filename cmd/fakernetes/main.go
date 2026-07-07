package main

import (
	"fmt"
	"os"

	"github.com/frailmink/fakernetes/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Printf("An error occurred when executing the command: %v", err)
		os.Exit(1)
	}
}