/*
Copyright © 2023 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"codeberg.org/tfkhdyt/prayermate/pkg/api"
	"codeberg.org/tfkhdyt/prayermate/pkg/stdout"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search all available locations by keyword",
	Long:  "Search all available locations by keyword",
	Run: func(_ *cobra.Command, args []string) {
		for _, keyword := range args {
			locations, err := api.SearchLocation(keyword)
			cobra.CheckErr(err)

			stdout.PrintLocationTable(locations)
		}
	},
}

func init() {
	locationCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
