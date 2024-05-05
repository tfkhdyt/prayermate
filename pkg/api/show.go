package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"codeberg.org/tfkhdyt/prayermate/entity"
	"codeberg.org/tfkhdyt/prayermate/pkg/fetch"
)

func ShowPrayerTimes(locationID string) (*entity.PrayerTimes, error) {
	year, month, day := time.Now().Date()
	url := fmt.Sprintf("https://api.myquran.com/v2/sholat/jadwal/%s/%d/%d/%d", locationID, year, int(month), day)

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	var result struct {
		Status bool `json:"status"`
		Data   struct {
			Jadwal entity.PrayerTimes `json:"jadwal"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if !result.Status {
		return nil, errors.New("failed to show prayer times")
	}

	return &result.Data.Jadwal, nil
}
