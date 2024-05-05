/*
Copyright © 2023 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"codeberg.org/tfkhdyt/prayermate/pkg/api"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show currently selected location",
	Long:  `Show currently selected location`,
	Run: func(_ *cobra.Command, _ []string) {
		locationID := viper.GetString("location.id")

		location, err := api.SearchLocationByID(locationID)
		cobra.CheckErr(err)

		fmt.Println("ID:", location.ID)
		fmt.Println("Name:", location.Name)
	},
}

func init() {
	locationCmd.AddCommand(currentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
