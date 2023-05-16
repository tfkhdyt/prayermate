/*
Copyright © 2023 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cmd

import (
	"codeberg.org/tfkhdyt/prayermate/pkg/api"
	"codeberg.org/tfkhdyt/prayermate/pkg/stdout"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search [location name...]",
	Short: "Search location by name",
	Long:  `Search location by name`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range args {
			locations, err := api.SearchLocation(name)
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
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
