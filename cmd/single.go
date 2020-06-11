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
	"github.com/currycan/scanner/core"

	"github.com/spf13/cobra"
)

var (
	name    string
	project string
	host    string
	port    int
)

// singleCmd represents the single command
var singleCmd = &cobra.Command{
	Use:   "single",
	Short: "check the connectivity of single server with command line",
	Long: `check the connectivity of single server with command line. 
For example:

scanner single --host 10.0.0.1 --port 6379 --name redis --project test
`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Scan(host, port, project, name)
	},
}

func init() {
	rootCmd.AddCommand(singleCmd)

	singleCmd.PersistentFlags().StringVarP(&name, "name", "", "", "the name of the server")
	singleCmd.PersistentFlags().StringVarP(&project, "project", "", "", "the project of the server")
	singleCmd.PersistentFlags().StringVarP(&host, "host", "", "", "the IP or domain name of the server")
	singleCmd.PersistentFlags().IntVarP(&port, "port", "", 0, "the port of the server")
}
