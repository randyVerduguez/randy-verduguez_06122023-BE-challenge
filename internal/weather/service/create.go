package service

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/internal/weather/model"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/pkg/erru"
)

type CreateParams struct {
	Name       string    `valid:"required"`
	Region     string    `valid:"required"`
	Country    string    `valid:"required"`
	TempC      float32   `valid:"required"`
	TempF      float32   `valid:"required"`
	FeelsLikeC float32   `valid:"required"`
	FeelsLikeF float32   `valid:"required"`
	WindMph    float32   `valid:"required"`
	WindKph    float32   `valid:"required"`
	Humidity   int       `valid:"required"`
	CreatedOn  time.Time `valid:"required"`
}

func (service Service) Create(ctx context.Context, params CreateParams) (model.Weather, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return model.Weather{}, erru.ErrArgument{Wrapped: err}
	}

	tx, err := service.repo.Db.BeginTxx(ctx, nil)

	if err != nil {
		return model.Weather{}, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.Weather{
		Name:       params.Name,
		Region:     params.Region,
		Country:    params.Country,
		TempC:      params.TempC,
		TempF:      params.TempF,
		FeelsLikeC: params.FeelsLikeC,
		FeelsLikeF: params.FeelsLikeF,
		WindMph:    params.WindMph,
		WindKph:    params.WindKph,
		Humidity:   params.Humidity,
		CreatedOn:  time.Now().UTC(),
	}

	err = service.repo.Create(ctx, &entity)

	if err != nil {
		return model.Weather{}, err
	}

	return entity, tx.Commit()
}
