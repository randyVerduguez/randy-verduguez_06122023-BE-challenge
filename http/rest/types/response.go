package types

import "time"

type Response struct {
	Name       string    `json:"name"`
	Region     string    `json:"region"`
	Country    string    `json:"country"`
	TempC      float32   `json:"temp_c"`
	TempF      float32   `json:"temp_f"`
	FeelsLikeC float32   `json:"feelsLike_c"`
	FeelsLikeF float32   `json:"feelsLike_f"`
	WindMph    float32   `json:"wind_mph"`
	WindKph    float32   `json:"wind_kph"`
	Humidity   int       `json:"humidity"`
	CreatedOn  time.Time `json:"created_on"`
}
