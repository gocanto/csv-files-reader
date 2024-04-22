package repository

import (
	"fmt"
	"github.com/gocanto/csv-files-reader/api/db"
	"github.com/gocanto/csv-files-reader/api/entity"
	"github.com/gocanto/csv-files-reader/api/handler"
	"github.com/gocanto/csv-files-reader/api/http"
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

func (receiver TradesRepository) InsertFromCSV(file handler.CSVFile) ([]entity.Trade, error) {
	query, err := receiver.conn.DB.Prepare(
		"INSERT INTO trades (unix, symbol, open, high, low, close) VALUES (?, ?, ?, ?, ?, ?)",
	)

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

func (receiver TradesRepository) Query(seed entity.Trade, pagination http.Pagination) ([]entity.Trade, error) {
	filters := entity.ParseTradesFiltersFrom(seed)
	baseQuery := "SELECT * FROM trades WHERE 1 = 1"
	args := make([]any, 0)

	for key, val := range filters {
		baseQuery += fmt.Sprintf(" AND %s = ?", key) // fuc()
		args = append(args, val)
	}

	baseQuery += fmt.Sprintf(" LIMIT %s OFFSET %s", pagination.Limit, pagination.Offset)

	rows, err := receiver.DB.DB.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var trades []entity.Trade

	for rows.Next() {
		var current entity.Trade

		if err := rows.Scan(
			&current.ID,
			&current.Unix,
			&current.Symbol,
			&current.Open,
			&current.High,
			&current.Low,
			&current.Close,
		); err != nil {
			return nil, err
		}

		trades = append(trades, current)
	}

	return trades, nil
}
