package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mytodo",
	Short: "MyTodo is a simple todo list app in cli",
	Run: func(cmd *cobra.Command, args []string) {
		if val, _ := cmd.Flags().GetBool("verbose"); val {
			log.Println("Verbose mode")
		}

		log.Println("Good Job")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
