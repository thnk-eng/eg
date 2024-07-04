package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitcli",
	Short: "Git CLI is a tool to manage git repositories",
	Long:  `A tool to manage git repositories with options to clone, configure, pull, and push.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func confirmAction(prompt string) bool {
	var response string
	fmt.Printf("%s (y/n): ", prompt)
	_, err := fmt.Scanln(&response)
	if err != nil {
		fmt.Println("Invalid input. Please enter y or n.")
		return confirmAction(prompt)
	}

	switch response {
	case "y", "Y":
		return true
	case "n", "N":
		return false
	default:
		fmt.Println("Please answer yes (y) or no (n).")
		return confirmAction(prompt)
	}
}
