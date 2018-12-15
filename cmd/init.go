package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dsimidzija/golang-goto/gotorepo"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [<repo-path>]",
	Short: "Initialize a goto repository",
	Long: `Creates a new empty alias repository in the provided path,
or in the default location ($HOME/.goto) if none is provided`,
	Run: func(cmd *cobra.Command, args []string) {
		var repo string
		var err error

		if len(args) == 1 {
			repo, err = gotorepo.GetRepoPath(args[0])
		} else {
			repo, err = gotorepo.GetRepoPath(viper.GetString("REPO_ROOT"))
		}
		DieIf(err)

		valid, err := gotorepo.IsValidRepo(repo)
		DieIf(err)

		if valid == true {
			fmt.Println(fmt.Sprintf("%s is already a goto repository", repo))
			os.Exit(1)
		}

		err = gotorepo.Init(repo)
		DieIf(err)
		fmt.Println(fmt.Sprintf("Initialized new goto repo at %s", repo))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
