package controller

import (
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

func (controller TradesController) Upload(w http.ResponseWriter, r *http.Request) {
	response := apiHttp.MakeResponse(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum
	if err != nil {
		http.Error(w, "Max form size is 10MB", http.StatusBadRequest)
		return
	}

	file, err := handler.MakeCSVFileFrom("file", r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := controller.repository.InsertFromCSV(entity.Trade{}, file)
	if err != nil {
		log.Fatal("Error with insert", err)
	}

	_ = response.Ok(output)
}
