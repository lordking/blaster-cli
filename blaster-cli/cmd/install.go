// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"log"
	"os/exec"

	blaster "github.com/lordking/blaster-cli"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install blaster-cli in your system.",
	Long: `Install will create .blaster in your $HOME directory, and create a
config file for blaster in .blaster.

If you don't succeed beacause of too long time, you can install again.

	`,
	Run: func(cmd *cobra.Command, args []string) {

		f, err := exec.LookPath("git")
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Printf("Found git in: %s", f)

		if config.Base == "" {
			config, err = blaster.CreateDefaultConfig()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		log.Printf("Found config directory at: %s", config.Base)

		repo = blaster.NewRepo(&config)
		if err := repo.Install(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(installCmd)
}
