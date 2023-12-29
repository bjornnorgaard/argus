package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dependabotCmd represents the dependabot command
var dependabotCmd = &cobra.Command{
	Use:     "dependabot",
	Aliases: []string{"bot"},
	Short:   "Handles dependabot pull requests",
	Long:    "Approves pull requests by dependabot that are passing status checks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dependabot called")
	},
}

func init() {
	githubCmd.AddCommand(dependabotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dependabotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dependabotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
