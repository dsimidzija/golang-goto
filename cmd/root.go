package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dsimidzija/golang-goto/gotorepo"
	"github.com/dsimidzija/golang-goto/gotorepo/alias"
)

var currentRepo string

var rootCmd = &cobra.Command{
	Use: "goto <alias>",
	Short: "Jump to specified dir alias",
	Args: cobra.ExactArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, args[] string) {
		if cmd.IsAvailableCommand() && cmd.Use != initCmd.Use {
			repo, err := gotorepo.GetRepoPath(viper.GetString("REPO_ROOT"))
			DieIf(err)

			valid, err := gotorepo.IsValidRepo(repo)
			DieIf(err)

			if valid != true {
				defaultPath, err := gotorepo.GetDefaultRepoPath()
				DieIf(err)
				var msg string
				if repo == defaultPath {
					msg = fmt.Sprintf("Directory %s is not a valid goto repository, run `goto init`", repo)
				} else {
					msg = fmt.Sprintf("Directory %s is not a valid goto repository, run `goto init %[1]s`", repo)
				}
				fmt.Println(msg)
				os.Exit(1)
			} else {
				currentRepo = repo
			}
		}
	},
	Run: func(cmd *cobra.Command, args[] string) {
		alias, err := alias.Load(currentRepo, args[0])
		DieIf(err)

		fmt.Println(alias.Data.Target)
		if cmd.PersistentFlags().Changed("signal") {
			signal, err := cmd.PersistentFlags().GetInt("signal")
			DieIf(err)
			os.Exit(signal)
		}
	},
}

// Execute runs the default cobra command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().IntP("signal", "s", 0, "Exit status code to let caller know it's safe to cd the output")
	rootCmd.PersistentFlags().MarkHidden("signal")
}

func initConfig() {
	viper.SetEnvPrefix("GOTO")
	viper.BindEnv("REPO_ROOT")

	defaultPath, err := gotorepo.GetDefaultRepoPath()
	DieIf(err)

	viper.SetDefault("REPO_ROOT", defaultPath)
	viper.AutomaticEnv()
}
