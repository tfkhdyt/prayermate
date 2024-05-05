package api

import (
	"encoding/json"

	"codeberg.org/tfkhdyt/prayermate/entity"
	"codeberg.org/tfkhdyt/prayermate/pkg/fetch"
)

func ListLocations() ([]entity.Location, error) {
	url := "https://api.myquran.com/v2/sholat/kota/semua"

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	var data struct {
		Locations []entity.Location `json:"data"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data.Locations, nil
}
