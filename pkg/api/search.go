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

type SearchLocationByIDDto struct {
	Status bool            `json:"status"`
	Data   entity.Location `json:"data"`
}

func SearchLocation(keyword string) ([]entity.Location, error) {
	url := fmt.Sprintf(
		"https://api.myquran.com/v2/sholat/kota/cari/%s",
		keyword,
	)

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	data := new(SearchLocationDto)

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data.Data) == 0 {
		return nil, errors.New("location is not found")
	}

	return data.Data, nil
}

func SearchLocationByID(id string) (*entity.Location, error) {
	url := fmt.Sprintf(
		"https://api.myquran.com/v2/sholat/kota/%s",
		id,
	)

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	data := new(SearchLocationByIDDto)

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if !data.Status {
		return nil, errors.New("location is not found")
	}

	return &data.Data, nil
}
