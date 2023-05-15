package alarm

import (
	"fmt"
	"log"
	"os"

	"github.com/gen2brain/beeep"
)

var shareDir = "/usr/share/prayermate/"

func notify(prayerType string, time string) {
	var icon string
	if _, err := os.Stat(shareDir + "assets/img/mosque.png"); os.IsNotExist(err) {
		icon = "assets/img/mosque.png"
	} else {
		icon = shareDir + "assets/img/mosque.png"
	}

	if err := beeep.Notify(prayerType+" Prayer", fmt.Sprintf("%s is the time of %s Prayer", time, prayerType), icon); err != nil {
		log.Fatalln("Error:", err.Error())
	}
}
