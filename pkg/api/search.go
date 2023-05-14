package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"codeberg.org/tfkhdyt/prayermate/entity"
	"codeberg.org/tfkhdyt/prayermate/pkg/fetch"
)

type SearchLocationDto struct {
	Status bool              `json:"status"`
	Data   []entity.Location `json:"data"`
}

func SearchLocation(name string) ([]entity.Location, error) {
	url := fmt.Sprintf("https://api.myquran.com/v1/sholat/kota/cari/%s", name)

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	var result SearchLocationDto
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if !result.Status {
		return nil, errors.New("Failed to search location")
	}

	return result.Data, nil
}
