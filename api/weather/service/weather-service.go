package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/api/weather/types/current"
)

func GetCurrentWeather(r current.CurrentWeatherRequest, jsonResBody *current.WeatherCurrentResponse) {
	baseUrl, _ := url.Parse("https://api.weatherapi.com/v1/current.json")
	params := url.Values{}

	params.Add("q", r.Q)
	params.Add("key", r.Key)
	baseUrl.RawQuery = params.Encode()

	httpResp, http_err := http.Get(baseUrl.String())

	if http_err != nil {
		error_msg := fmt.Errorf("Error: API returned %d with error %s", httpResp.StatusCode, http_err)
		fmt.Println(error_msg)
		return
	}

	resbody, read_err := io.ReadAll(httpResp.Body)

	defer httpResp.Body.Close()

	if read_err != nil {
		error_msg := fmt.Errorf("Error: Unable to read http response body: %s", read_err)
		fmt.Println(error_msg)
		return
	}

	json_err := json.Unmarshal(resbody, jsonResBody)

	if json_err != nil {
		error_msg := fmt.Errorf("Error: Unable to parse http response body: %s", read_err)
		fmt.Println(error_msg)
		return
	}
}
