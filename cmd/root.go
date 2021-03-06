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

//import "C"
import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/currycan/scanner/core"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debug bool

var cfgFile string
var c core.Config
var svc core.Services
var s core.Service
var configChange = make(chan int, 1)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scanner",
	Short: "scan the server port ",
	Long: `A tool that check the connectivity of specific server or servers,
you can set a yaml config file for multiple services or use command line for a single server.
For example:
	scanner -c /path/to/example/scanner.yaml
	scanner --host 10.0.0.1 --port 6379 --name redis --project test
If you want to open the debug mode,you can add '--debug true in the end of command
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("args: " + strings.Join(args, " "))
		// multi
		config, _ := cmd.Flags().GetString("config")
		// single
		lineName, _ := cmd.Flags().GetString("name")
		lineProject, _ := cmd.Flags().GetString("project")
		lineHost, _ := cmd.Flags().GetString("host")
		linePort, _ := cmd.Flags().GetInt("port")
		logrus.Debugf("Name: %s, Host: %s, Port: %d", lineName, lineHost, linePort)

		if config != "" {
			for _, s = range svc.Services {
				logrus.Debugf("Name: %s, Host: %s, Port: %d", s.Name, s.Host, s.Port)
				core.Scan(s.Host, s.Port, s.Project, s.Name)
			}
			// 等待配置改变, 然后重启
			c.WatchConfig(configChange)
			for ; ; {
				<-configChange
				for _, s = range svc.Services {
					logrus.Debugf("Name: %s, Host: %s, Port: %d", s.Name, s.Host, s.Port)
					core.Scan(s.Host, s.Port, s.Project, s.Name)
				}
			}
		} else if lineName != "" && lineProject != "" && lineHost != "" && linePort != 0 {
			core.Scan(lineHost, linePort, lineProject, lineName)
		} else {
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
	// For multiple services with config file
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file path")
	// For single service with command line
	rootCmd.PersistentFlags().StringVarP(&core.Name, "name", "", "", "the name of the server")
	rootCmd.PersistentFlags().StringVarP(&core.Project, "project", "", "", "the project of the server")
	rootCmd.PersistentFlags().StringVarP(&core.Host, "host", "", "", "the IP or domain name of the server")
	rootCmd.PersistentFlags().IntVarP(&core.Port, "port", "", 0, "the port of the server")
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

// initConfig 读取配置文件和环境变量
func initConfig() {
	var err error
	c.Name = cfgFile
	if err, svc = c.UnmarshalStruct(); err != nil {
		fmt.Errorf("%s", err.Error())
		os.Exit(1)
	}
}
