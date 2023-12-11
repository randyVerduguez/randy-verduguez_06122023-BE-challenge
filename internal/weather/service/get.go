package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	weatherService "github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/api/weather/service"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/api/weather/types/current"
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
		request := createAPIRequest(params)
		responseBody := callWeatherAPI(request)
		newEntry, create_err := s.Create(ctx, CreateParams{
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

		if create_err != nil {
			error_msg := fmt.Errorf("Error: unable to create DB entry: %s", create_err)
			return model.Weather{}, erru.ErrArgument{errors.New(error_msg.Error())}
		}

		return newEntry, create_err
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
	days := int64(now.Sub(entryDate).Hours() / 24)

	if days > 0 {
		request := createAPIRequest(params)
		responseBody := callWeatherAPI(request)
		updatedEntry, update_err := s.Update(ctx, currentEntry, UpdateParams{
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

		if update_err != nil {
			error_msg := fmt.Errorf("Unable to update DB entry: %s", update_err)
			return model.Weather{}, erru.ErrArgument{errors.New(error_msg.Error())}
		}

		return updatedEntry, update_err
	}

	return currentEntry, nil
}

func callWeatherAPI(request current.CurrentWeatherRequest) *current.WeatherCurrentResponse {
	var response current.WeatherCurrentResponse

	weatherService.GetCurrentWeather(request, &response)

	return &response
}

func createAPIRequest(params types.Request) current.CurrentWeatherRequest {
	return current.CurrentWeatherRequest{
		Q:   fmt.Sprintf("%s,%s,%s", params.Name, params.Region, params.Country),
		Key: "5c78b517e79c415a937172805230612",
	}
}
