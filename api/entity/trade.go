package entity

import (
	"github.com/gocanto/csv-files-reader/api/support"
	"strings"
)

var AllowedInFiltering = []string{"symbol", "unix"}

type Trade struct {
	ID     string  `json:"id"`
	Unix   string  `json:"unix"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
}

func ParseTradesFiltersFrom(trade Trade) map[string]string {
	filter := make(map[string]string)
	allowed := AllowedInFiltering

	if seed := len(strings.TrimSpace(trade.Symbol)); seed > 0 && support.ContainsString(allowed, "symbol") {
		filter["symbol"] = trade.Symbol
	}

	if seed := len(strings.TrimSpace(trade.Unix)); seed > 0 && support.ContainsString(allowed, "unix") {
		filter["unix"] = trade.Unix
	}

	return filter
}
