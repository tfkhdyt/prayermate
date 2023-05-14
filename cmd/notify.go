/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"codeberg.org/tfkhdyt/prayermate/entity"
	"codeberg.org/tfkhdyt/prayermate/pkg/api"
	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

var locationID string

// notifyCmd represents the notify command
var notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Show notifications when it's time to pray",
	Long:  `Show notifications when it's time to pray`,
	Run: func(cmd *cobra.Command, args []string) {
		prayerTimes, err := api.ShowPrayerTimes(locationID)
		if err != nil {
			log.Fatalf("Error: %v\n", err.Error())
		}

		for {
			checkTime(prayerTimes)
			time.Sleep(1 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notifyCmd.PersistentFlags().String("foo", "", "A help for foo")
	notifyCmd.PersistentFlags().StringVarP(&locationID, "locationId", "l", "1301", "Set location")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkTime(prayerTimes *entity.PrayerTimes) {
	hourNow, minuteNow := time.Now().Hour(), time.Now().Minute()

	// subuh
	var hourSubuh, minuteSubuh int
	if _, err := fmt.Sscanf(prayerTimes.Data.Schedule.Subuh, "%d:%d", &hourSubuh, &minuteSubuh); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if hourNow == hourSubuh && minuteNow == minuteSubuh {
		notify("Subuh", prayerTimes.Data.Schedule.Subuh)
	}

	// dzuhur
	var hourDzuhur, minuteDzuhur int
	if _, err := fmt.Sscanf(prayerTimes.Data.Schedule.Dzuhur, "%d:%d", &hourDzuhur, &minuteDzuhur); err != nil {
		log.Fatalln("Error:", err.Error())
	}

	if hourNow == hourDzuhur && minuteNow == minuteDzuhur {
		notify("Dzuhur", prayerTimes.Data.Schedule.Dzuhur)
	}
}

func notify(prayerType string, time string) {
	if err := beeep.Notify(prayerType+"Prayer", fmt.Sprintf("%s is the time of Magrib Prayer", time), ""); err != nil {
		log.Fatalln("Error:", err.Error())
	}
}
