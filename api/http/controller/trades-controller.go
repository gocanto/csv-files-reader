package controller

import (
	"fmt"
	"net/http"
	"ohlc-price-data/api/entity"
	"ohlc-price-data/api/handler"
	apiHttp "ohlc-price-data/api/http"
	"ohlc-price-data/api/repository"
)

type TradesController struct {
	repository repository.TradesRepository
}

func (controller TradesController) Upload(w http.ResponseWriter, r *http.Request) {
	response := apiHttp.MakeResponse(w, r)

	if r.Method != http.MethodPost {
		_ = response.MethodNotAllowed()
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum
	if err != nil {
		_ = response.BadRequest("Invalid input. Max form size is 10MB")
		return
	}

	file, err := handler.MakeCSVFileFrom("file", r)
	if err != nil {
		_ = response.BadRequest(fmt.Sprintf("There was an issue [%s] while reading the file.", err))
		return
	}

	//@todo Add validation on duplicated entries.
	output, err := controller.repository.InsertFromCSV(entity.Trade{}, file)
	if err != nil {
		_ = response.ServerError(fmt.Sprintf("There was an issue [%s] inserting the parsed data.", err))

		return
	}

	_ = response.Ok(output)
}
