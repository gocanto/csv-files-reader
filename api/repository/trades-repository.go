package repository

import (
	"fmt"
	"ohlc-price-data/api/db"
	"ohlc-price-data/api/entity"
	"ohlc-price-data/api/handler"
	"strconv"
)

type TradesRepository struct {
	conn db.Connection
	DB   db.Connection
}

func MakeTradesRepositoryFrom(conn db.Connection) (TradesRepository, error) {
	return TradesRepository{
		conn: conn,
		DB:   conn,
	}, nil
}

func (receiver TradesRepository) InsertFromCSV(trade entity.Trade, file handler.CSVFile) ([]entity.Trade, error) {
	query, err := receiver.conn.DB.Prepare(trade.GetInsertSQL())

	if err != nil {
		return nil, err
	}

	var output []entity.Trade
	for key, val := range file.GetContent() {
		if key == 0 { //header
			continue
		}

		open, _ := strconv.ParseFloat(val[2], 64)
		high, _ := strconv.ParseFloat(val[3], 64)
		low, _ := strconv.ParseFloat(val[4], 64)
		closeVal, _ := strconv.ParseFloat(val[5], 64)

		output = append(output, entity.Trade{
			Unix:   val[0],
			Symbol: val[1],
			Open:   open,
			High:   high,
			Low:    low,
			Close:  closeVal,
		})

		_, err := query.Exec(val[0], val[1], val[2], val[3], val[4], val[5])

		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

func (receiver TradesRepository) Query(limit, offset int, filter map[string]interface{}) ([]entity.Trade, error) {
	baseQuery := "SELECT * FROM trades WHERE 1 = 1"

	args := make([]interface{}, 0)
	var conditions []string

	for field, value := range filter {
		switch field {
		case "symbol":
			conditions = append(conditions, fmt.Sprintf("%s = ?", "symbol"))
			args = append(args, value)
		case "open":
			conditions = append(conditions, fmt.Sprintf("%s = ?", "open"))
			args = append(args, value)
		case "high":
			conditions = append(conditions, fmt.Sprintf("%s = ?", "high"))
			args = append(args, value)
		}
	}

	if len(conditions) > 0 {
		for _, cond := range conditions {
			baseQuery += " AND " + cond
		}
	}

	baseQuery += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	fmt.Println(baseQuery, "-----", args, "-- conditions --", conditions)

	// Execute the query
	rows, err := receiver.conn.DB.Query(baseQuery, args...)
	if err != nil {
		fmt.Println("-- here --")
		return nil, err
	}
	//defer rows.Close()

	var trades []entity.Trade

	// Iterate through the result set
	for rows.Next() {
		var t entity.Trade
		if err := rows.Scan(&t.Unix, &t.Unix, &t.Symbol, &t.Open, &t.High, &t.Low, &t.Close); err != nil {
			fmt.Println("-- here 2 --")
			return nil, err
		}

		trades = append(trades, t)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		fmt.Println("-- here 3--")
		return nil, err
	}

	return trades, nil
}
