package entity

type PrayerTimes struct {
	Status bool `json:"status"`
	Data   struct {
		ID         string `json:"id"`
		Location   string `json:"lokasi"`
		Province   string `json:"daerah"`
		Coordinate struct {
			Latitude     float64 `json:"lat"`
			Longitude    float64 `json:"lon"`
			LatitudeStr  string  `json:"lintang"`
			LongitudeStr string  `json:"bujur"`
		} `json:"koordinat"`
		Schedule struct {
			Date   string `json:"tanggal"`
			Imsak  string `json:"imsak"`
			Subuh  string `json:"subuh"`
			Terbit string `json:"terbit"`
			Dhuha  string `json:"dhuha"`
			Dzuhur string `json:"dzuhur"`
			Ashar  string `json:"ashar"`
			Magrib string `json:"maghrib"`
			Isya   string `json:"isya"`
		} `json:"jadwal"`
	} `json:"data"`
}
