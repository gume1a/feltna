package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/gume1a/feltna/prettyprint"
	"github.com/spf13/cobra"
)

// List of repositories to clone
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

// init is a function that adds the clone command to the root command and adds the path flag to the clone command.
func init() {
	versionCmd.Flags().StringP("path", "p", "", "path to where the repositories should be cloned")
	rootCmd.AddCommand(versionCmd)
}

// versionCmd is a command that clones all gume1a repositories to the given path.
var versionCmd = &cobra.Command{
	Use:   "clone",
	Short: "command to clone all gume1a repositories",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")

		for _, repo := range repositories {
			// Print cloning repo with green color
			prettyprint.CPrintF(prettyprint.GREEN, "%s\n", "CLONE", repo)
			exec.
				Command("git", "clone", fmt.Sprintf("https://github.com/gume1a/%s.git", repo), filepath.Join(path, repo)).
				Run()
		}
	},
}
