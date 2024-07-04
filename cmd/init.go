package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize GitHub SSH and Git configuration",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter your GitHub email: ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)

		fmt.Print("Enter your GitHub username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)

		if confirmAction("Generate a new SSH key?") {
			fmt.Println("Generating a new SSH key...")
			cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-f", os.Getenv("HOME")+"/.ssh/id_ed25519", "-N", "")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return
			}
		}

		if confirmAction("Start the SSH agent?") {
			fmt.Println("Starting the SSH agent...")
			cmd := exec.Command("ssh-agent", "-s")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return
			}
		}

		if confirmAction("Add the SSH private key to the SSH agent?") {
			fmt.Println("Adding the SSH private key to the SSH agent...")
			cmd := exec.Command("ssh-add", os.Getenv("HOME")+"/.ssh/id_ed25519")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return
			}
		}

		if confirmAction("Configure SSH for GitHub?") {
			fmt.Println("Configuring SSH for GitHub...")
			sshConfig := `
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519
  IdentitiesOnly yes
`
			f, _ := os.OpenFile(os.Getenv("HOME")+"/.ssh/config", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {

				}
			}(f)
			_, err := f.WriteString(sshConfig)
			if err != nil {
				return
			}
		}

		if confirmAction("Configure Git?") {
			fmt.Println("Configuring Git...")
			cmd := exec.Command("git", "config", "--global", "user.name", username)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return
			}

			cmd = exec.Command("git", "config", "--global", "user.email", email)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				return
			}
		}

		if confirmAction("Display the SSH public key?") {
			fmt.Println("Displaying the SSH public key...")
			fmt.Println("Copy the following SSH public key and add it to your GitHub account:")
			cmd := exec.Command("cat", os.Getenv("HOME")+"/.ssh/id_ed25519.pub")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return
			}

			fmt.Println("\nInstructions to add the SSH key to your GitHub account:")
			fmt.Println("1. Go to https://github.com/settings/keys")
			fmt.Println("2. Click 'New SSH key' or 'Add SSH key'")
			fmt.Println("3. Paste the key below into the 'Key' field")
			fmt.Println("4. Click 'Add SSH key'")
		}

		if confirmAction("Test SSH connection to GitHub?") {
			fmt.Println("Testing SSH connection to GitHub...")
			cmd := exec.Command("ssh", "-T", "git@github.com")
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
	rootCmd.AddCommand(initCmd)
}
