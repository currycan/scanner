/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import "C"
import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scanner",
	Short: "scan the server port ",
	Long: `A tool that check the connectivity of specific server or servers,
you can set a yaml config file for multiple services or use command line for a single server.
The default yaml config file is $HOME/scanner.yaml. 
For example:
scanner multi (use default config: $HOME/scanner.yaml，should exist)
scanner multi -c /path/to/conf/scanner.yaml (specify config)
scanner single --host 10.0.0.1 --port 6379 --name redis --project test
If you want to open the debug mode,you can add '--debug true in the end of command
	`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
	},
}

func Execute() {
	subCmd, _, err := rootCmd.Find([]string{filepath.Base(os.Args[0])})
	if err == nil && subCmd.Name() != rootCmd.Name() {
		if len(os.Args) > 1 {
			rootCmd.SetArgs(append([]string{subCmd.Name()}, os.Args[1:]...))
		} else {
			rootCmd.SetArgs([]string{subCmd.Name()})
		}
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLog)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug mode")
}

func initLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
