package main

import (
	"os"

	"github.com/spf13/cobra"

	"interact_cli_test/internal/run"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run repo-scaffolding",
	Long:  `...`,
	Run:   runCmdFunc,
}

func runCmdFunc(cmd *cobra.Command, args []string) {
	if err := run.Execute(); err != nil {
		os.Exit(1)
	}
}
