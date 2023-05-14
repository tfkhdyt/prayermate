/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"codeberg.org/tfkhdyt/prayermate/pkg/api"
	"codeberg.org/tfkhdyt/prayermate/pkg/stdout"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available locations",
	Long:  "List all available locations",
	Run: func(cmd *cobra.Command, args []string) {
		locations, err := api.ListLocations()
		if err != nil {
			log.Fatalf("Error: %v\n", err.Error())
		}

		stdout.PrintLocationTable(locations)
	},
}

func init() {
	locationCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
