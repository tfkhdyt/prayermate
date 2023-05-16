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

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set [location id]",
	Short: "Set location to config file",
	Long:  `Set location to config file`,
	Run: func(cmd *cobra.Command, args []string) {
		locationID := args[0]

		locationDetail, err := api.GetLocationDetail(locationID)
		cobra.CheckErr(err)

		viper.Set("location.id", locationDetail.Data.ID)

		cobra.CheckErr(viper.WriteConfig())

		fmt.Println("Selected location:", locationDetail.Data.Location)
		fmt.Println("The selected location has been saved in the config file")
	},
}

func init() {
	locationCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
