package alarm

import (
	"strconv"
	"strings"
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

	splittedTime := strings.Split(timeStr, ":")
	hour, _ := strconv.Atoi(splittedTime[0])
	minute, _ := strconv.Atoi(splittedTime[1])

	if hourNow == hour && minuteNow == minute {
		notify(prayerType, timeStr)
	}
}
