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
	"os"
	"path"

	blaster "github.com/lordking/blaster-cli"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var config blaster.Config
var templates blaster.Templates
var repo blaster.Repo

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "blaster-cli",
	Short: "A simple CLI for scaffolding Blaster Projects.",
	Long: `The 'blaster-cli' is a CLI tool for scaffolding golang projects.
You can quickly create a web project or other projects.

  `,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blaster/config.yaml or .blaster/config.yaml)")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")          // name of config file (without extension)
	viper.AddConfigPath("$HOME/.blaster/") // adding home directory as first search path
	viper.AddConfigPath(".blaster/")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {

		form := blaster.ConfigForm{}
		if err := viper.Unmarshal(&form); err != nil {
			fmt.Printf("error: %s", err)
		}

		configDir := path.Dir(viper.ConfigFileUsed())
		config = blaster.GetConfig(form, configDir)

		repo = blaster.NewRepo(&config)
		templates = blaster.NewTemplates(&config, &repo)
	}

	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// })
}
