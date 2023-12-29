package cmd

import (
	"github.com/bjornnorgaard/argus/tools/github"

	"github.com/spf13/cobra"
)

// githubCmd represents the github command
var githubCmd = &cobra.Command{
	Use:     "github",
	Aliases: []string{"gh"},
	Short:   "Uses GitHub CLI to interact with repos",
	Long:    "Use one of the nested commands to actually do something",
	Run: func(cmd *cobra.Command, args []string) {
		github.Run()
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// githubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// githubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
