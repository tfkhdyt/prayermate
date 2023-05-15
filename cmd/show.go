/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"codeberg.org/tfkhdyt/prayermate/pkg/api"
	"codeberg.org/tfkhdyt/prayermate/pkg/stdout"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show [location id...]",
	Short: "Show prayer times for this day",
	Long:  `Show prayer times for this day`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			locationID := viper.GetString("locationId")
			showSchedule(locationID)
		} else {
			for _, locationID := range args {
				showSchedule(locationID)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showSchedule(locationID string) {
	prayerTimes, err := api.ShowPrayerTimes(locationID)
	if err != nil {
		log.Fatalf("Error: %v\n", err.Error())
	}

	fmt.Printf(`ID: %s
Location: %s
Province: %s
Coordinate:
  - Latitude: %s
  - Longitude: %s
Date: %s
`, prayerTimes.Data.ID, prayerTimes.Data.Location, prayerTimes.Data.Province, prayerTimes.Data.Coordinate.LatitudeStr, prayerTimes.Data.Coordinate.LongitudeStr, prayerTimes.Data.Schedule.Date)

	stdout.PrintScheduleTable(prayerTimes)
}
