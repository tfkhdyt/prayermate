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

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set [location id]",
	Short: "Set location to config file",
	Long:  `Set location to config file`,
	Run: func(_ *cobra.Command, args []string) {
		locationID := args[0]

		location, err := api.SearchLocationByID(locationID)
		cobra.CheckErr(err)

		viper.Set("location.id", location.ID)

		cobra.CheckErr(viper.WriteConfig())

		fmt.Println("Selected location:", location.Name)
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
