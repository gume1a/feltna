package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var repositories = []string{
	"feltna",
	"management-frontend",
	"dev-portal",
	"g1a-infrastructure",
	"management-service",
	"management-service-lib",
	"documentation",
	"g1a-node-api",
	"g1a-angular-frontend",
	"keycloak-extension",
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "git",
	Short: "Command to run git operations over the organization's repositories.",
	Run: func(cmd *cobra.Command, args []string) {

		for _, repo := range repositories {
			exec.
				Command("git", "clone", fmt.Sprintf("https://github.com/gume1a/%s.git", repo)).
				Run()
		}
	},
}
