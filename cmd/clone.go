package cmd

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	githubapi "github.com/google/go-github/v50/github"
	"github.com/gume1a/feltna/internal/middleware"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"log"
	"os"
	"path/filepath"
)

// cloneCmd is a command that clones all Gume1a repositories to the given path.
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "command to clone all gume1a repositories",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")

		middleware.GithubAuthentication([]string{"repo", "read:org"}, func(token *oauth2.Token, client *githubapi.Client) {
			repositories, _, err := client.
				Repositories.
				ListByOrg(context.Background(), "gume1a", nil)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s cloning %s repositories into %s\n", color.BlueString("STARTED"), color.HiMagentaString("gume1a"), color.BlueString(path))
			for _, repo := range repositories {

				_, err := git.PlainClone(filepath.Join(path, repo.GetName()), false, &git.CloneOptions{
					URL:      fmt.Sprintf("https://oauth2:%s@github.com/gume1a/%s.git", token.AccessToken, repo.GetName()),
					Progress: os.Stdout,
				})

				if err != nil {
					if err == git.ErrRepositoryAlreadyExists {
						fmt.Printf("%s %s\n", color.YellowString("EXISTS"), repo.GetName())
						continue
					}
					fmt.Printf("%s %s: %v\n", color.RedString("ERROR"), repo.GetName(), err)
					continue
				}

				fmt.Printf("%s %s\n", color.GreenString("CLONED"), repo.GetName())
			}

		})
	},
}

// init is a function that adds the clone command to the root command and adds the path flag to the clone command.
func init() {
	cloneCmd.Flags().StringP("path", "p", "./gume1a", "path to where the repositories should be cloned")
	rootCmd.AddCommand(cloneCmd)
}
