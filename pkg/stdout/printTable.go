package stdout

import (
	"os"

	"codeberg.org/tfkhdyt/prayermate/entity"
	"github.com/olekukonko/tablewriter"
)

func PrintTable(locations []entity.Location) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Location"})

	for _, location := range locations {
		table.Append([]string{location.ID, location.Location})
	}
	table.Render() // Send output
}
