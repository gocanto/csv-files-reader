package controller

import "github.com/gocanto/csv-files-reader/api/repository"

func MakeTradesController(repository repository.TradesRepository) (TradesController, error) {
	return TradesController{
		repository: repository,
	}, nil
}
