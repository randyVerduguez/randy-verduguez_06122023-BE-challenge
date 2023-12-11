package current

type CurrentWeatherRequest struct {
	Q   string `json:"q"`
	Key string `json:"key"`
}
