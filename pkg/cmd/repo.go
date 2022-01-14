package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tmiddlet2666/ghstats/pkg/config"
	"github.com/tmiddlet2666/ghstats/pkg/utils"
)

// getRepo implements the 'get rep' command
var getRepo = &cobra.Command{
	Use:   "repo",
	Short: "Display the repository details for a user and repository",
	Long:  `The 'get repo' command displays the repository details a GitHub repository.`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		var (
			err         error
			repoDetails config.Repository
		)

		repoDetails, err = utils.GetRepoDetails(userName, repo)
		if err != nil {
			return err
		}

		cmd.Println()
		cmd.Printf("Repository URL: %s\n", utils.GetRepositoryURL(userName, repo))
		cmd.Printf("Name:           %s\n", repoDetails.Name)
		cmd.Printf("Full Name:      %s\n", repoDetails.FullName)
		cmd.Printf("Description:    %s\n", repoDetails.Description)
		cmd.Printf("Language:       %s\n", repoDetails.Language)
		cmd.Printf("Stars:          %d\n", repoDetails.Stars)
		cmd.Printf("Subscribers:    %d\n", repoDetails.Subscribers)
		cmd.Printf("Forks:          %d\n", repoDetails.Forks)

		return nil
	},
}
