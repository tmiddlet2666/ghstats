package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/tmiddlet2666/ghstats/pkg/config"
	"github.com/tmiddlet2666/ghstats/pkg/utils"
)

var (
	tag      string
	fileName string
)

// getDownloadsCmd implements the 'get downloads' command
var getDownloadsCmd = &cobra.Command{
	Use:   "downloads",
	Short: "Display the downloads for assets for a user and repository",
	Long:  `The 'get downloads' command displays the downloads for asssets for a GitHub repository.`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {

		var (
			err            error
			releases       []config.Release
			headerFmt      = color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt      = color.New(color.FgYellow).SprintfFunc()
			found          = false
			totalDownloads int64
		)

		releases, err = utils.GetReleases(userName, repo)
		if err != nil {
			return err
		}

		// validate the tag
		if tag != "all" {
			for _, value := range releases {
				if value.TagName == tag {
					found = true
					break
				}
			}

			if !found {
				return fmt.Errorf("unable to find tag %s for user: %s and repository %s", tag, userName, repo)
			}
		}

		tbl := table.New("TAG", "Name", "Created", "        Size", " Downloads")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, value := range releases {
			var tagValue = value.TagName
			if tag != "all" && tag != value.TagName {
				continue
			}

			for _, asset := range value.Assets {
				if fileName == "" || fileName == asset.Name {
					tbl.AddRow(tagValue, asset.Name, value.CreatedAt,
						utils.FormatFileSize(asset.Size), utils.FormatLargeInteger(asset.DownloadCount))
					tagValue = ""
					totalDownloads += asset.DownloadCount
				}
			}
		}

		tbl.AddRow("", "", "", "", "----------")
		tbl.AddRow("", "", "", "TOTAL", utils.FormatLargeInteger(totalDownloads))
		cmd.Printf("\nRepository: %s\n", utils.GetRepositoryURL(userName, repo))
		tbl.Print()

		return nil
	},
}

func init() {
	getDownloadsCmd.Flags().StringVarP(&tag, "tag", "t", "all", "tag to show downloads for")
	getDownloadsCmd.Flags().StringVarP(&fileName, "file-name", "f", "", "file name to show downloads for")
}
