package models

import "time"

//type CityAqiDetails struct {
//	CityName string `json:"cityName" db:"city_name"`
//	Aqi      int    `json:"aqi" db:"city_aqi"`
//}
const (
	City string = "delhi"
)

type CityAqiDetails struct {
	Message string        `json:"message"`
	Lat     float64       `json:"lat"`
	Lng     float64       `json:"lng"`
	Data    []Combination `json:"data"`
}
type Combination struct {
	Ozone       Ozone       `json:"ozone"`
	Co2         Co2         `json:"co2"`
	Ch4         Ch4         `json:"ch4"`
	WaterVapour WaterVapour `json:"waterVapour"`
}

type Ozone struct {
	Value       string    `json:"value" db:"ozone_value"`
	LastUpdated time.Time `json:"lastUpdated" db:"ozone_last_updated"`
	Units       string    `json:"units" db:"ozone_units"`
}

type Co2 struct {
	Value       string    `json:"value" db:"co2_value"`
	LastUpdated time.Time `json:"lastUpdated" db:"co2_last_updated"`
	Units       string    `json:"units" db:"co2_units"`
}

type Ch4 struct {
	Value       string    `json:"value" db:"ch4_value"`
	LastUpdated time.Time `json:"lastUpdated" db:"ch4_last_updated"`
	Units       string    `json:"units" db:"ch4_units"`
}

type WaterVapour struct {
	Value       int       `json:"value"`
	LastUpdated time.Time `json:"lastUpdated"`
	Units       string    `json:"units"`
}
