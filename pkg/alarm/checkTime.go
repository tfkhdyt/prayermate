package alarm

import (
	"fmt"
	"time"

	"codeberg.org/tfkhdyt/prayermate/entity"
)

func CheckTime(prayerTimes *entity.PrayerTimes) {
	day := time.Now().Weekday().String()

	parseTimeAndNotify(prayerTimes.Subuh, "Subuh")
	parseTimeAndNotify(prayerTimes.Ashar, "Ashar")
	parseTimeAndNotify(prayerTimes.Magrib, "Magrib")
	parseTimeAndNotify(prayerTimes.Isya, "Isya")
	if day == "Friday" {
		parseTimeAndNotify(prayerTimes.Dzuhur, "Friday")
	} else {
		parseTimeAndNotify(prayerTimes.Dzuhur, "Dzuhur")
	}
}

func parseTimeAndNotify(timeStr string, prayerType string) {
	hourNow, minuteNow := time.Now().Hour(), time.Now().Minute()

	var hour int
	var minute int
	fmt.Sscanf(timeStr, "%02d:%02d", &hour, &minute)

	if hourNow == hour && minuteNow == minute {
		notify(prayerType, timeStr)
	}
}
