package handler

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"net/http"
)

type CSVFile struct {
	file    multipart.File
	handler *multipart.FileHeader
	content [][]string
}

func MakeCSVFileFrom(key string, r *http.Request) (CSVFile, error) {
	file, handler, err := r.FormFile(key)
	defer file.Close()

	if err != nil {
		return CSVFile{}, err
	}

	if handler.Header.Get("Content-Type") != "text/csv" {
		return CSVFile{}, fmt.Errorf("invalid file format. Please upload a CSV file")
	}

	content, err := csv.NewReader(file).ReadAll()

	if err != nil {
		return CSVFile{}, err
	}

	return CSVFile{
		file:    file,
		handler: handler,
		content: content,
	}, nil
}

func (receiver CSVFile) GetContent() [][]string {
	return receiver.content
}
