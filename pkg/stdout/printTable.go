package stdout

import (
	"os"

	"codeberg.org/tfkhdyt/prayermate/entity"
	"github.com/olekukonko/tablewriter"
)

func PrintLocationTable(locations []entity.Location) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Location"})

	for _, location := range locations {
		table.Append([]string{location.ID, location.Location})
	}
	table.Render() // Send output
}

func PrintScheduleTable(prayerTimes *entity.PrayerTimes) {
	table := tablewriter.NewWriter(os.Stdout)

	table.Append([]string{"Imsak", prayerTimes.Data.Schedule.Imsak})
	table.Append([]string{"Subuh", prayerTimes.Data.Schedule.Subuh})
	table.Append([]string{"Terbit", prayerTimes.Data.Schedule.Terbit})
	table.Append([]string{"Dhuha", prayerTimes.Data.Schedule.Dhuha})
	table.Append([]string{"Dzuhur", prayerTimes.Data.Schedule.Dzuhur})
	table.Append([]string{"Ashar", prayerTimes.Data.Schedule.Ashar})
	table.Append([]string{"Magrib", prayerTimes.Data.Schedule.Magrib})
	table.Append([]string{"Isya", prayerTimes.Data.Schedule.Isya})

	table.Render() // Send output
}
