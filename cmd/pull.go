package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull changes",
	Run: func(cmd *cobra.Command, args []string) {
		if confirmAction("Pull changes?") {
			cmd := exec.Command("git", "pull", "origin", "main")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
