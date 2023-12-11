package model

import "time"

type Weather struct {
	Id         int       `db:"id"`
	Name       string    `db:"name"`
	Region     string    `db:"region"`
	Country    string    `db:"country"`
	TempC      float32   `db:"temp_c"`
	TempF      float32   `db:"temp_f"`
	FeelsLikeC float32   `db:"feelslike_c"`
	FeelsLikeF float32   `db:"feelslike_f"`
	WindMph    float32   `db:"wind_mph"`
	WindKph    float32   `db:"wind_kph"`
	Humidity   int       `db:"humidity"`
	CreatedOn  time.Time `db:"created_on"`
}
