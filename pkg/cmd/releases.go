package cmd

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/tmiddlet2666/ghstats/pkg/config"
	"github.com/tmiddlet2666/ghstats/pkg/utils"
)

// getTags implements the 'get tags' command
var getTags = &cobra.Command{
	Use:   "releases",
	Short: "Display the releases for a user and repository",
	Long:  `The 'get releases' command displays the releases for a GitHub repository.`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		var (
			err       error
			releases  []config.Release
			headerFmt = color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt = color.New(color.FgYellow).SprintfFunc()
		)

		releases, err = utils.GetReleases(userName, repo)
		if err != nil {
			return err
		}
		tbl := table.New("TAG", "Name", "Pre-release?", "Published", "Assets")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, value := range releases {
			tbl.AddRow(value.TagName, value.Name, value.PreRelease, value.PublishedAt, len(value.Assets))
		}
		cmd.Printf("\nRepository: %s\n", utils.GetRepositoryURL(userName, repo))
		tbl.Print()

		return nil
	},
}
