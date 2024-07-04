package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure Git repository (Add remote)",
	Run: func(cmd *cobra.Command, args []string) {
		if confirmAction("Configure Git repository (Add remote)?") {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter the SSH URL of the remote repository: ")
			repoURL, _ := reader.ReadString('\n')
			repoURL = strings.TrimSpace(repoURL)
			cmd := exec.Command("git", "remote", "add", "origin", repoURL)
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
	rootCmd.AddCommand(configureCmd)
}
