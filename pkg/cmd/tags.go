package cmd

import (
	"github.com/spf13/cobra"
)

// getTags implements the 'get tags' command
var getTags = &cobra.Command{
	Use:   "tags",
	Short: "Display the list tags for a repository",
	Long:  `The 'get tags' command displays the list of tags for a GitHub repository.`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		cmd.Println(
			"get tags " + userName + ", " + repo,
		)
		return nil
	},
}
