package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConfigDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func Connect(config ConfigDB) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	fmt.Println(dataSourceName)
	db, err := sqlx.Connect("postgres", dataSourceName)

	return db, err
}
