package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
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

	//fmt.Println(file, "handler", content, "---------", err)

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

	//for record := content {
	//	nixTime, _ := strconv.ParseInt(record[0], 10, 64)
	//	open, _ := strconv.ParseFloat(record[2], 64)
	//	high, _ := strconv.ParseFloat(record[3], 64)
	//	low, _ := strconv.ParseFloat(record[4], 64)
	//	closeRow, _ := strconv.ParseFloat(record[5], 64)
	//}
	//
	//for {
	//	record := content
	//	if err != nil {
	//		fmt.Println("reader error", err, "data:", record)
	//		break
	//	}
	//
	//	unixTime, _ := strconv.ParseInt(record[0], 10, 64)
	//	open, _ := strconv.ParseFloat(record[2], 64)
	//	high, _ := strconv.ParseFloat(record[3], 64)
	//	low, _ := strconv.ParseFloat(record[4], 64)
	//	closeRow, _ := strconv.ParseFloat(record[5], 64)
	//	data = append(data, StockData{
	//		UnixTime: unixTime,
	//		Symbol:   record[1],
	//		Open:     open,
	//		High:     high,
	//		Low:      low,
	//		Close:    closeRow,
	//	})
	//}
	//
	//// Connect to MySQL database
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "pass", "localhost", 3306, "trades")
	//fmt.Println("dsn:", dsn)
	//db, err := sql.Open("mysql", "gocanto:gocanto@tcp(db:3306)/trades")
	////db, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	log.Fatal("Error connecting to database:", err)
	//}
	//defer db.Close()
	////db.SetConnMaxLifetime(time.Minute * 3)
	////db.SetMaxOpenConns(10)
	////db.SetMaxIdleConns(10)
	//
	//// Insert data into MySQL table
	//stmt, err := db.Prepare("INSERT INTO trades (unix, symbol, open, high, low, close) VALUES (?, ?, ?, ?, ?, ?)")
	//if err != nil {
	//	log.Fatal("Error preparing statement:", err)
	//}
	//defer stmt.Close()
	//
	//for _, d := range data {
	//	fmt.Println("data:", d, "---")
	//
	//	_, err := stmt.Exec(d.UnixTime, d.Symbol, d.Open, d.High, d.Low, d.Close)
	//	if err != nil {
	//		log.Fatal("Error inserting row:", err)
	//	}
	//}
	//
	//fmt.Println("Data inserted successfully")
}

type StockData struct {
	UnixTime int64
	Symbol   string
	Open     float64
	High     float64
	Low      float64
	Close    float64
}
