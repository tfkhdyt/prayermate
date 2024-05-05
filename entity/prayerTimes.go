package entity

type PrayerTimes struct {
	Date   string `json:"tanggal"`
	Imsak  string `json:"imsak"`
	Subuh  string `json:"subuh"`
	Terbit string `json:"terbit"`
	Dhuha  string `json:"dhuha"`
	Dzuhur string `json:"dzuhur"`
	Ashar  string `json:"ashar"`
	Magrib string `json:"maghrib"`
	Isya   string `json:"isya"`
}
