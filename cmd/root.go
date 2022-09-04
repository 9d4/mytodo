package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mytodo",
	Short: "MyTodo is a simple todo list app in cli",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(lsCmd, args)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
