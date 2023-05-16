package alarm

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
)

func notify(prayerType string, time string) {
	shareDir := getShareDir()

	var icon string
	if _, err := os.Stat(filepath.Join(shareDir, "assets/img/mosque.png")); os.IsNotExist(err) {
		icon = "assets/img/mosque.png"
	} else {
		icon = filepath.Join(shareDir, "assets/img/mosque.png")
	}

	cobra.CheckErr(
		beeep.Notify(
			prayerType+" Prayer",
			fmt.Sprintf("%s is the time of %s Prayer", time, prayerType),
			icon,
		),
	)
}

func getShareDir() string {
	if _, err := os.Stat("/usr/local/share/prayermate/"); os.IsNotExist(err) {
		return "/usr/share/prayermate/"
	}
	return "/usr/local/share/prayermate/"
}
