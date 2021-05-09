/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Display and clear logs",
	Long:  "Display and clear logs",
}

var displayLogsCmd = &cobra.Command{
	Use:   "display",
	Short: "Display all logs",
	RunE: func(cmd *cobra.Command, args []string) error {
		logs, err := strg.Log().GetAll()
		if err != nil {
			return err
		}
		cmd.Println("Logs")
		for _, log := range logs {
			cmd.Println(`------------------------------`)
			cmd.Println("ID: ", log.ID)
			cmd.Println("Message: ", log.Message)
			cmd.Println("Created at: ", log.CreatedAt)
		}

		return nil
	},
}

var clearLogsCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all logs",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := strg.Log().DeleteAll()
		if err != nil {
			return err
		}

		cmd.Println("Logs cleared!")
		return nil
	},
}

func init() {
	logCmd.AddCommand(displayLogsCmd)
	logCmd.AddCommand(clearLogsCmd)

	rootCmd.AddCommand(logCmd)
}
