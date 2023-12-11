package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/http/rest/types"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/internal/weather/model"
	"github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/pkg/db"
)

type Repository struct {
	Db *sqlx.DB
}

type Params struct {
	Name    string
	City    string
	Country string
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{Db: db}
}

func (repository Repository) Find(ctx context.Context, params types.Request) (model.Weather, error) {
	entity := model.Weather{}
	query := fmt.Sprintf("SELECT * FROM weather WHERE name = $1 AND region = $2 AND country = $3;")
	err := repository.Db.GetContext(ctx, &entity, query, params.Name, params.Region, params.Country)

	return entity, db.HandleError(err)
}

func (repository Repository) Create(ctx context.Context, entity *model.Weather) error {
	query := `INSERT INTO weather (
		name, region, country, temp_c, temp_f, feelslike_c, feelslike_f, wind_mph, wind_kph, humidity, created_on
	) VALUES (
		:name, :region, :country, :temp_c, :temp_f, :feelslike_c, :feelslike_f, :wind_mph, :wind_kph, :humidity, :created_on
	);`

	rows, err := repository.Db.NamedQueryContext(ctx, query, entity)

	if err != nil {
		return db.HandleError(err)
	}

	for rows.Next() {
		err = rows.StructScan(entity)

		if err != nil {
			return db.HandleError(err)
		}
	}

	return err
}

func (respository Repository) Update(ctx context.Context, entity model.Weather) error {
	query := `
		UPDATE weather 
		SET temp_c = :temp_c,
			temp_f = :temp_f,
			wind_mph = :wind_mph,
			wind_kph = :wind_kph,
			humidity = :humidity,
			created_on = :created_on
		WHERE id = :id;`

	_, err := respository.Db.NamedExecContext(ctx, query, entity)

	if err != nil {
		fmt.Println(err)
	}

	return db.HandleError(err)
}
