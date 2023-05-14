package api

import (
	"encoding/json"

	"codeberg.org/tfkhdyt/prayermate/entity"
	"codeberg.org/tfkhdyt/prayermate/pkg/fetch"
)

func ListLocations() ([]entity.Location, error) {
	url := "https://api.myquran.com/v1/sholat/kota/semua"

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	var locations []entity.Location

	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}
