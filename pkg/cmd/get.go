package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "display one or many resources",
	Long:  `The 'get' command displays one or more resources.`,
}
