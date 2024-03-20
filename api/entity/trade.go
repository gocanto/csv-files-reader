package entity

type Trade struct {
	Unix   string  `json:"unix"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
}

func (receiver Trade) GetInsertSQL() string {
	return "INSERT INTO trades (unix, symbol, open, high, low, close) VALUES (?, ?, ?, ?, ?, ?)"
}
