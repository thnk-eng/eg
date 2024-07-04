package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes",
	Run: func(cmd *cobra.Command, args []string) {
		if confirmAction("Push changes?") {
			reader := bufio.NewReader(os.Stdin)
			cmd := exec.Command("git", "add", ".")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return
			}

			fmt.Print("Enter the commit message: ")
			commitMessage, _ := reader.ReadString('\n')
			commitMessage = strings.TrimSpace(commitMessage)
			cmd = exec.Command("git", "commit", "-m", commitMessage)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return
			}

			cmd = exec.Command("git", "push", "origin", "main")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
