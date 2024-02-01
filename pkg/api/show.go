package api

import (
	"encoding/json"
	"fmt"
	"time"

	"codeberg.org/tfkhdyt/prayermate/entity"
	"codeberg.org/tfkhdyt/prayermate/pkg/fetch"
)

func ShowPrayerTimes(locationID string) (*entity.PrayerTimes, error) {
	year, month, day := time.Now().Date()
	url := fmt.Sprintf("https://cdn.statically.io/gh/lakuapik/jadwalsholatorg/master/adzan/%s/%d/%02d.json", locationID, year, int(month))

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	var result []entity.PrayerTimes
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result[day-1], nil
}
