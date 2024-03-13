package di

import (
	"github.com/onemariovalentino/fetch-webpage/src/app/fetch/repositories"
	"github.com/onemariovalentino/fetch-webpage/src/app/fetch/usecases"
)

type (
	Dependency struct {
		FetchUsecase usecases.FetchUsecaseInterface
	}
)

func New() *Dependency {
	fetchRepository := repositories.New("files/json/fetch_data.json")
	fetchUsecase := usecases.New(fetchRepository)

	return &Dependency{
		FetchUsecase: fetchUsecase,
	}
}
