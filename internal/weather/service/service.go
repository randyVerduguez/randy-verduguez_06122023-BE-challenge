package service

import "github.com/randyVerduguez/randy-verduguez_06122023-BE-challenge/internal/weather/repository"

type Service struct {
	repo repository.Repository
}

func NewService(repository repository.Repository) Service {
	return Service{
		repo: repository,
	}
}
