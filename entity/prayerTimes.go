package entity

type PrayerTimes struct {
	Date   string `json:"tanggal"`
	Imsak  string `json:"imsyak"`
	Subuh  string `json:"shubuh"`
	Terbit string `json:"terbit"`
	Dhuha  string `json:"dhuha"`
	Dzuhur string `json:"dzuhur"`
	Ashar  string `json:"ashr"`
	Magrib string `json:"magrib"`
	Isya   string `json:"isya"`
}
