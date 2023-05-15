package api

import (
	"encoding/json"
	"errors"

	"codeberg.org/tfkhdyt/prayermate/pkg/fetch"
)

type VerifyDto struct {
	Status bool `json:"status"`
	Data   struct {
		ID       string `json:"id"`
		Location string `json:"lokasi"`
	} `json:"data"`
}

func GetLocationDetail(locationID string) (*VerifyDto, error) {
	url := "https://api.myquran.com/v1/sholat/kota/id/" + locationID

	body, err := fetch.GET(url)
	if err != nil {
		return nil, err
	}

	var status VerifyDto

	if err := json.Unmarshal(body, &status); err != nil {
		return nil, err
	}

	if !status.Status {
		return nil, errors.New("Location id is not found")
	}

	return &status, nil
}
