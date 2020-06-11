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

import (
	"fmt"
	"os"

	"github.com/currycan/scanner/core"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool

type Outbound struct {
	Name    string
	Project string
	Host    string
	Port    int
}

type Inbound struct {
	Name    string
	Project string
	Host    string
	Port    int
}

type Services struct {
	Outbounds []Outbound
	Inbounds  []Outbound
}

var conf Services

// multiCmd represents the multi command
var multiCmd = &cobra.Command{
	Use:   "multi",
	Short: "check the connectivity of multi servers",
	Long: `check the connectivity of multi servers with yaml config file.
The default yaml config file is $HOME/scanner.yaml. 
For example:

scanner multi (use default config: $HOME/scanner.yaml，should exist)
scanner multi -c /path/to/conf/scanner.yaml (specify config)
`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, sOut := range conf.Outbounds {
			logrus.Debugf("Name: %s, Host: %s, Port: %d", sOut.Name, sOut.Host, sOut.Port)
			core.Scan(sOut.Host, sOut.Port, sOut.Project, sOut.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(multiCmd)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/scanner.yaml)")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, _ := homedir.Dir()
		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(home)
		//viper.AddConfigPath("conf")
		viper.SetConfigName("scanner")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		fmt.Println("You can run 'scanner multi --help' to get help")
		os.Exit(1)
	}

	err := viper.Unmarshal(&conf)
	if err != nil {
		panic("Unable to unmarshal config")
	}
}
