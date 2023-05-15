package api

import (
	"encoding/json"
	"errors"

	"codeberg.org/tfkhdyt/prayermate/pkg/fetch"
)

type VerifyDto struct {
	Status bool `json:"status"`
}

func VerifyLocationIDAvailability(locationId string) error {
	url := "https://api.myquran.com/v1/sholat/kota/id/" + locationId

	body, err := fetch.GET(url)
	if err != nil {
		return err
	}

	var status VerifyDto

	if err := json.Unmarshal(body, &status); err != nil {
		return err
	}

	if !status.Status {
		return errors.New("Location id is not found")
	}

	return nil
}
