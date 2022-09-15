package cmd

import (
	"github.com/spf13/cobra"
	"runtime"
)

// versionCmd implements the 'get versionCmd' command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version information",
	Long:  `The 'get version' shows the version information.`,
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("GitHub Stats\nVersion:      %s\nDate:         %s\n"+
			"Commit:       %s\nOS:           %s\nArchitecture: %s\n",
			Version, Date, Commit, runtime.GOOS, runtime.GOARCH)
		return nil
	},
}
