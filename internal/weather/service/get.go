package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	weatherService "github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/api/weather/service"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/api/weather/types/current"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/configs"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/http/rest/types"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/internal/weather/model"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/pkg/db"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/pkg/erru"
)

func (s Service) Get(ctx context.Context, params types.Request) (model.Weather, error) {
	currentEntry, err := s.repo.Find(ctx, params)

	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		request, err := createAPIRequest(params)

		if err != nil {
			return model.Weather{}, err
		}

		responseBody := callWeatherAPI(request)
		newEntry, err := s.Create(ctx, CreateParams{
			Name:       strings.ToLower(responseBody.Location.Name),
			Region:     strings.ToLower(responseBody.Location.Region),
			Country:    strings.ToLower(responseBody.Location.Country),
			TempC:      responseBody.Current.TempC,
			TempF:      responseBody.Current.TempF,
			FeelsLikeC: responseBody.Current.FeelsLikeC,
			FeelsLikeF: responseBody.Current.FeelsLikeF,
			WindMph:    responseBody.Current.WindMph,
			WindKph:    responseBody.Current.WindKph,
			Humidity:   responseBody.Current.Humidity,
			CreatedOn:  time.Now().UTC(),
		})

		if err != nil {
			error_msg := fmt.Errorf("Error: unable to create DB entry: %s", err)
			return model.Weather{}, erru.ErrArgument{errors.New(error_msg.Error())}
		}

		return newEntry, err
	default:
		return model.Weather{}, err
	}

	now := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
		time.Now().Nanosecond(),
		time.UTC,
	)
	entryDate := time.Date(
		currentEntry.CreatedOn.Year(),
		currentEntry.CreatedOn.Month(),
		currentEntry.CreatedOn.Day(),
		currentEntry.CreatedOn.Hour(),
		currentEntry.CreatedOn.Minute(),
		currentEntry.CreatedOn.Second(),
		currentEntry.CreatedOn.Nanosecond(),
		time.UTC,
	)

	// if current entry is older than today, then update it
	days := now.Day() - entryDate.Day()

	if days > 0 {
		request, err := createAPIRequest(params)

		if err != nil {
			return model.Weather{}, err
		}

		responseBody := callWeatherAPI(request)
		updatedEntry, err := s.Update(ctx, currentEntry, UpdateParams{
			Id:         currentEntry.Id,
			TempC:      responseBody.Current.TempC,
			TempF:      responseBody.Current.TempF,
			FeelsLikeC: responseBody.Current.FeelsLikeC,
			FeelsLikeF: responseBody.Current.FeelsLikeF,
			WindMph:    responseBody.Current.WindMph,
			WindKph:    responseBody.Current.WindKph,
			Humidity:   responseBody.Current.Humidity,
			CreatedOn:  time.Now().UTC(),
		})

		if err != nil {
			error_msg := fmt.Errorf("Unable to update DB entry: %s", err)
			return model.Weather{}, erru.ErrArgument{errors.New(error_msg.Error())}
		}

		return updatedEntry, err
	}

	return currentEntry, nil
}

func callWeatherAPI(request current.CurrentWeatherRequest) *current.WeatherCurrentResponse {
	var response current.WeatherCurrentResponse

	weatherService.GetCurrentWeather(request, &response)

	return &response
}

func createAPIRequest(params types.Request) (current.CurrentWeatherRequest, error) {
	config, err := configs.NewParsedConfig()

	if err != nil {
		error := fmt.Errorf("Unable to update DB entry: %s", err)
		return current.CurrentWeatherRequest{}, error
	}

	return current.CurrentWeatherRequest{
		Q:   fmt.Sprintf("%s,%s,%s", params.Name, params.Region, params.Country),
		Key: config.ApiKey,
	}, err
}
