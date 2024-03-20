package repository

import (
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
