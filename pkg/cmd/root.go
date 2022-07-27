package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/tmiddlet2666/ghstats/pkg/constants"
	"os"
)

var (
	Version string
	Commit  string
	Date    string
	rootCmd = createRootCommand()

	userName string
	repo     string
)

// createRootCommand creates the root command off which all others are places
func createRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:          "ghstats",
		Short:        "Get GitHub Stats",
		SilenceUsage: true,
		Long:         `This command allows you to retrieve various stats about GitHub Repositories.`,
	}
	return root
}

// Execute run the root command
func Execute(version string, date string, commit string) {
	Version = version
	Date = date
	Commit = commit
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// initialize commands
	initializeGlobalFlags()

	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getReleasesCmd)
	getCmd.AddCommand(getDownloadsCmd)
	getCmd.AddCommand(getRepo)

	rootCmd.AddCommand(versionCmd)

	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)
}

func initializeGlobalFlags() {
	// setup global flags
	rootCmd.PersistentFlags().StringVarP(&userName, constants.UsernameFlag, "u", "", "GitHub username")
	rootCmd.PersistentFlags().StringVarP(&repo, constants.RepositoryFlag, "r", "", "GitHub repository")
}

func validateUserAndRepo() error {
	if userName == "" || repo == "" {
		return errors.New("you must provide username and repository")
	}
	return nil
}
