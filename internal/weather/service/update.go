package service

import (
	"context"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/internal/weather/model"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/pkg/erru"
)

type UpdateParams struct {
	Id         int       `valid:"required"`
	TempC      float32   `valid:"required"`
	TempF      float32   `valid:"required"`
	FeelsLikeC float32   `valid:"required"`
	FeelsLikeF float32   `valid:"required"`
	WindMph    float32   `valid:"required"`
	WindKph    float32   `valid:"required"`
	Humidity   int       `valid:"required"`
	CreatedOn  time.Time `valid:"required"`
}

func (s Service) Update(ctx context.Context, entity model.Weather, params UpdateParams) (model.Weather, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return model.Weather{}, erru.ErrArgument{Wrapped: err}
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return model.Weather{}, err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// update entity
	entity.TempC = params.TempC
	entity.TempF = params.TempF
	entity.FeelsLikeC = params.FeelsLikeC
	entity.FeelsLikeF = params.FeelsLikeF
	entity.WindMph = params.WindMph
	entity.WindKph = params.WindKph
	entity.Humidity = params.Humidity
	entity.CreatedOn = params.CreatedOn

	err = s.repo.Update(ctx, entity)

	if err != nil {
		return model.Weather{}, err
	}

	return entity, tx.Commit()
}
