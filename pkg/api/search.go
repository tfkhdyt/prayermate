package api

import (
	"errors"

	"codeberg.org/tfkhdyt/prayermate/entity"
)

type SearchLocationDto struct {
	Status bool              `json:"status"`
	Data   []entity.Location `json:"data"`
}

func SearchLocation(name string) ([]string, error) {
	locations, err := ListLocations()
	if err != nil {
		return nil, err
	}

	for _, location := range locations {
		if location == name {
			return []string{location}, nil
		}
	}

	return nil, errors.New("location is not found")
}
