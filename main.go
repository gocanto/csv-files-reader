package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/upload/csv", uploadCSVHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadCSVHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := csv.NewReader(file).ReadAll()

	//// Ensure the uploaded file is a CSV file
	if handler.Header.Get("Content-Type") != "text/csv" {
		http.Error(w, "Invalid file format. Please upload a CSV file", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(content)
	if err != nil {
		fmt.Println("--------->", err)
		return
	}

	// Connect to MySQL database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "gocanto", "gocanto", "db", 3306, "trades")
	fmt.Println("dsn:", dsn, "current: ", "gocanto:gocanto@tcp(db:3306)/trades")
	db, err := sql.Open("mysql", "gocanto:gocanto@tcp(db:3306)/trades")

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	var output []StockData

	stmt, err := db.Prepare("INSERT INTO trades (unix, symbol, open, high, low, close) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error with statement")
	}
	defer stmt.Close()

	for key, val := range content {
		if key == 0 { //header
			continue
		}

		open, _ := strconv.ParseFloat(val[2], 64)
		high, _ := strconv.ParseFloat(val[3], 64)
		low, _ := strconv.ParseFloat(val[4], 64)
		closeVal, _ := strconv.ParseFloat(val[5], 64)

		output = append(output, StockData{
			unix:   val[0],
			Symbol: val[1],
			Open:   open,
			High:   high,
			Low:    low,
			Close:  closeVal,
		})

		fmt.Println("key:", key, "val:", val)

		_, err := stmt.Exec(val[0], val[1], val[2], val[3], val[4], val[5])
		if err != nil {
			fmt.Println("error inserting")
			return
		}
	}

	fmt.Println("Output: ", output)
	fmt.Println("Data inserted successfully")
}

type StockData struct {
	unix   string
	Symbol string
	Open   float64
	High   float64
	Low    float64
	Close  float64
}
