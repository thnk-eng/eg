package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if confirmAction("Clone a repository?") {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter the SSH URL of the repository to clone: ")
			repoURL, _ := reader.ReadString('\n')
			repoURL = strings.TrimSpace(repoURL)
			cmd := exec.Command("git", "clone", repoURL)
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
	rootCmd.AddCommand(cloneCmd)
}
