/*
Copyright © 2023 Taufik Hidayat <tfkhdyt@proton.me>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"codeberg.org/tfkhdyt/prayermate/pkg/alarm"
	"codeberg.org/tfkhdyt/prayermate/pkg/api"
)

var (
	locationID string
	isSilent   bool
)

// notifyCmd represents the notify command
var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Show notifications when it's time to pray",
	Long:  `Show notifications when it's time to pray`,
	Run: func(_ *cobra.Command, _ []string) {
		if locationID == "" {
			locationID = viper.GetString("location.id")
		}
		prayerTimes, err := api.ShowPrayerTimes(locationID)
		cobra.CheckErr(err)

		if !isSilent {
			fmt.Println("PrayerMate notify is running...")
			fmt.Println("Selected location:", locationID)
		}

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
	notifyCmd.PersistentFlags().BoolVarP(&isSilent, "silent", "s", false, "Hide the output")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
