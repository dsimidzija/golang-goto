// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/dsimidzija/golang-goto/gotorepo"
	"github.com/dsimidzija/golang-goto/gotorepo/alias"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <alias> <target-path>",
	Short: "Add a new goto alias",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("add needs two arguments, alias and target path")
		}

		regex := `(?i)^[a-z0-9\-_]+$`
		re := regexp.MustCompile(regex)
		segments := strings.Split(args[0], string(os.PathSeparator))
		for _, segment := range segments {
			if re.MatchString(segment) != true {
				fmt.Println(segment)
				return fmt.Errorf("%s is not a valid alias name, path elements must match \"%s\"", args[0], regex)
			}
		}

		target := filepath.Clean(args[1])
		var err error
		if filepath.IsAbs(args[1]) != true {
			target, err = filepath.Abs(target)
			DieIf(err)
		}

		valid, err := gotorepo.IsValidTarget(target)
		DieIf(err)
		if valid != true {
			return fmt.Errorf("%s is not a valid goto target", target)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		target, err := filepath.Abs(filepath.Clean(args[1]))
		DieIf(err)
		a := alias.New(args[0], target)
		exists, err := a.Exists(currentRepo)
		DieIf(err)

		if exists == true {
			force, err := cmd.Flags().GetBool("force");
			DieIf(err)

			if force == false {
				fmt.Println(fmt.Sprintf("Alias %s already exists, add --force flag to overwrite", a.Alias))
				os.Exit(1)
			}
		}

		err = a.Save(currentRepo)
		DieIf(err)

		fmt.Println(fmt.Sprintf("Added: %s => %s", a.Alias, a.Data.Target))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("force", "f", false, "Force overwrite if alias already exists")
}
