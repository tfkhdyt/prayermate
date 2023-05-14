package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"codeberg.org/tfkhdyt/prayermate/pkg/fetch"
)

type Coordinate struct {
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"lon"`
	LatitudeStr  string  `json:"lintang"`
	LongitudeStr string  `json:"bujur"`
}

type Schedule struct {
	Date   string `json:"tanggal"`
	Imsak  string `json:"imsak"`
	Subuh  string `json:"subuh"`
	Terbit string `json:"terbit"`
	Dhuha  string `json:"dhuha"`
	Dzuhur string `json:"dzuhur"`
	Ashar  string `json:"ashar"`
	Magrib string `json:"magrib"`
	Isya   string `json:"isya"`
}

type PrayerTimesData struct {
	ID         string     `json:"id"`
	Location   string     `json:"lokasi"`
	Province   string     `json:"daerah"`
	Coordinate Coordinate `json:"koordinat"`
	Schedule   Schedule   `json:"jadwal"`
}

type ShowPrayerTimesDto struct {
	Status bool            `json:"status"`
	Data   PrayerTimesData `json:"data"`
}

func ShowPrayerTimes(locationID string) (*ShowPrayerTimesDto, error) {
	year, month, day := time.Now().Date()
	url := fmt.Sprintf("https://api.myquran.com/v1/sholat/jadwal/%s/%d/%d/%d", locationID, year, int(month), day)

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	var result ShowPrayerTimesDto
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if !result.Status {
		return nil, errors.New("Failed to show prayer times")
	}

	return &result, nil
}
