package validator

import (
	"ohlc-price-data/api/entity"
	"strings"
)

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ParseTradesFiltersFrom(trade entity.Trade) map[string]string {
	filter := make(map[string]string)
	allowed := entity.AllowedInFiltering

	if seed := len(strings.TrimSpace(trade.Symbol)); seed > 0 && ContainsString(allowed, "symbol") {
		filter["symbol"] = trade.Symbol
	}

	if seed := len(strings.TrimSpace(trade.Unix)); seed > 0 && ContainsString(allowed, "unix") {
		filter["unix"] = trade.Unix
	}

	return filter
}
