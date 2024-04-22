package http

import (
	"net/url"
	"strings"
)

const DefaultLimit = "2"
const DefaultOffset = "0"

type Pagination struct {
	Limit  string //string for simplicity, but it should be int
	Offset string //string for simplicity, but it should be int
}

func MakeDefaultPaginationFrom(values url.Values) Pagination {
	pagination := make(map[string]string)
	pagination["limit"] = DefaultLimit
	pagination["offset"] = DefaultOffset

	if limit, ok := values["limit"]; ok {
		pagination["limit"] = strings.Join(limit, "")
	}

	if offset, ok := values["offset"]; ok {
		pagination["offset"] = strings.Join(offset, "")
	}

	return Pagination{
		Limit:  pagination["limit"],
		Offset: pagination["offset"],
	}
}
