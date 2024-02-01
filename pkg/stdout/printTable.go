package stdout

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	"codeberg.org/tfkhdyt/prayermate/entity"
)

func PrintLocationTable(locations []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Location"})

	for _, location := range locations {
		table.Append([]string{location})
	}
	table.Render() // Send output
}

func PrintScheduleTable(prayerTimes *entity.PrayerTimes) {
	table := tablewriter.NewWriter(os.Stdout)

	table.Append([]string{"Imsak", prayerTimes.Imsak})
	table.Append([]string{"Subuh", prayerTimes.Subuh})
	table.Append([]string{"Terbit", prayerTimes.Terbit})
	table.Append([]string{"Dhuha", prayerTimes.Dhuha})
	table.Append([]string{"Dzuhur", prayerTimes.Dzuhur})
	table.Append([]string{"Ashar", prayerTimes.Ashar})
	table.Append([]string{"Magrib", prayerTimes.Magrib})
	table.Append([]string{"Isya", prayerTimes.Isya})

	table.Render() // Send output
	fmt.Println()
}
