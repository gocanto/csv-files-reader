package controller

import "ohlc-price-data/api/repository"

func MakeTradesController(repository repository.TradesRepository) (TradesController, error) {
	return TradesController{
		repository: repository,
	}, nil
}
