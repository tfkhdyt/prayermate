/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"codeberg.org/tfkhdyt/prayermate/pkg/alarm"
	"codeberg.org/tfkhdyt/prayermate/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var locationID string

// notifyCmd represents the notify command
var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Show notifications when it's time to pray",
	Long:  `Show notifications when it's time to pray`,
	Run: func(cmd *cobra.Command, args []string) {
		if locationID == "" {
			locationID = viper.GetString("location.id")
		}
		prayerTimes, err := api.ShowPrayerTimes(locationID)
		cobra.CheckErr(err)

		fmt.Println("PrayerMate notify is running...")
		fmt.Println("Selected location:", prayerTimes.Data.Location)

		for {
			alarm.CheckTime(prayerTimes)

			time.Sleep(1 * time.Minute)
		}
	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	notifyCmd.PersistentFlags().StringVarP(&locationID, "locationId", "l", "", "Set location")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
