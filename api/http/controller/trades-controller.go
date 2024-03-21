package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ohlc-price-data/api/entity"
	"ohlc-price-data/api/handler"
	apiHttp "ohlc-price-data/api/http"
	"ohlc-price-data/api/repository"
)

type TradesController struct {
	repository repository.TradesRepository
}

func (controller TradesController) Query(w http.ResponseWriter, r *http.Request) {
	response := apiHttp.MakeResponse(w, r)

	if r.Method != http.MethodGet {
		_ = response.MethodNotAllowed("GET")
		return
	}

	var body entity.Trade
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		_ = response.BadRequest("The given payload is invalid.")
		return
	}

	trades, err := controller.repository.Query(
		body,
		entity.MakeDefaultPaginationFrom(r.URL.Query()),
	)

	if err != nil {
		_ = response.ServerError("There was an issue while fetching the data.")
		log.Fatal(fmt.Sprintf("Invalid query: [%s]", err))
		return
	}

	_ = response.Ok(trades)
}

func (controller TradesController) Upload(w http.ResponseWriter, r *http.Request) {
	response := apiHttp.MakeResponse(w, r)

	if r.Method != http.MethodPost {
		_ = response.MethodNotAllowed("POST")
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum
	if err != nil {
		_ = response.BadRequest("Invalid input. Max form size is 10MB")
		return
	}

	//@todo it can be multiple type of files, so the handler can be swap?
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

	//@todo better format?
	_ = response.Ok(output)
}
