package forecast

type WeatherForcastResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Latitude       float32 `json:"lat"`
	Longitude      float32 `json:"lon"`
	TzId           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      int64   `json:"localtime"`
}

type Current struct {
	LastUpdatedEpoch int64   `json:"last_updated_epoch"`
	LastUpdated      int64   `json:"last_updated"`
	TempC            float32 `json:"temp_c"`
	TempF            float32 `json:"temp_f"`
	IsDay            int     `json:"is_day"`
	WindMph          float32 `json:"wind_mph"`
	WindKph          float32 `json:"wind_kph"`
	WindDegree       int     `json:"wind_degree"`
	WindDirection    string  `json:"wind_dir"`
	PressureMb       float32 `json:"pressure_mb"`
	PressureIn       float32 `json:"pressure_in"`
	PrecipMm         float32 `json:"precip_mm"`
	PrecipIn         float32 `json:"precip_in"`
	Humidity         int     `json:"humidity"`
	Cloud            int     `json:"cloud"`
	FeelsLikeC       float32 `json:"feelslike_c"`
	FeelsLikeF       float32 `json:"feelslike_f"`
	VisKm            float32 `json:"vis_km"`
	VisMiles         float32 `json:"vis_miles"`
	UV               float32 `json:"uv"`
	GustMph          float32 `json:"gust_mph"`
	GustKph          float32 `json:"gust_kph"`
}
