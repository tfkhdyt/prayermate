/*
Copyright © 2023 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"codeberg.org/tfkhdyt/prayermate/pkg/api"
	"codeberg.org/tfkhdyt/prayermate/pkg/stdout"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available locations",
	Long:  "List all available locations",
	Run: func(_ *cobra.Command, _ []string) {
		locations, err := api.ListLocations()
		cobra.CheckErr(err)

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
