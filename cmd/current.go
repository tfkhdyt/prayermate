/*
Copyright © 2023 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cmd

import (
	"fmt"

	"codeberg.org/tfkhdyt/prayermate/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show currently selected location",
	Long:  `Show currently selected location`,
	Run: func(cmd *cobra.Command, args []string) {
		locationID := viper.GetString("location.id")
		locationDetail, err := api.GetLocationDetail(locationID)
		cobra.CheckErr(err)

		fmt.Println("ID:", locationDetail.Data.ID)
		fmt.Println("Location:", locationDetail.Data.Location)
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
