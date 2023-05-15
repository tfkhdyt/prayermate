/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"codeberg.org/tfkhdyt/prayermate/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Show currently selected location",
	Long:  `Show currently selected location`,
	Run: func(cmd *cobra.Command, args []string) {
		locationId := viper.GetString("location.id")
		locationDetail, err := api.GetLocationDetail(locationId)
		cobra.CheckErr(err)

		fmt.Println("ID:", locationDetail.Data.ID)
		fmt.Println("Location:", locationDetail.Data.Location)
	},
}

func init() {
	locationCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
