package forecast

type WeatherForcastRequest struct {
	Q    string `json:"q"`
	Days int    `json:"days"`
	Key  string `json:"key"`
}
