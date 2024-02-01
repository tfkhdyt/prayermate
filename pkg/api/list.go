package api

import (
	"encoding/json"

	"codeberg.org/tfkhdyt/prayermate/pkg/fetch"
)

func ListLocations() ([]string, error) {
	url := "https://raw.githubusercontent.com/lakuapik/jadwalsholatorg/master/kota.json"

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	var locations []string

	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}
