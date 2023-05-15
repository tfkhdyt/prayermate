package alarm

import (
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

var shareDir = "/usr/local/share/prayermate/"

func notify(prayerType string, time string) {
	var icon string
	if _, err := os.Stat(shareDir + "assets/img/mosque.png"); os.IsNotExist(err) {
		icon = "assets/img/mosque.png"
	} else {
		icon = shareDir + "assets/img/mosque.png"
	}

	cobra.CheckErr(
		beeep.Notify(
			prayerType+" Prayer",
			fmt.Sprintf("%s is the time of %s Prayer", time, prayerType),
			icon,
		),
	)
}
