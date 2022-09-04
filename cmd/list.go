package cmd

import "github.com/spf13/cobra"

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "show todo list items",
	Run:   func(cmd *cobra.Command, args []string) {},
}
