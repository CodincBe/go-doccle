package main

import (
	"fmt"
	"strings"
)

type RequestOptions struct {
	Lang     string
	Order    string
	Page     int
	PageSize int
	Sort     string
}

func DefaultRequestOptions() RequestOptions {
	options := RequestOptions{
		Lang:     "en",
		Order:    "DESC",
		Page:     1,
		PageSize: 50,
		Sort:     "Date",
	}

	return options
}

func (options RequestOptions) toQuery() string {
	query := []string{}
	if len(options.Lang) > 0 {
		query = append(query, fmt.Sprintf("lang=%s", options.Lang))
	}
	if options.Order == "DESC" || options.Order == "ASC" {
		query = append(query, fmt.Sprintf("order=%s", options.Order))
	}
	if options.Page > 0 {
		query = append(query, fmt.Sprintf("page=%d", options.Page))
	}
	if options.PageSize > 0 {
		query = append(query, fmt.Sprintf("pageSize=%d", options.PageSize))
	}
	// Haven't mapped the others atm
	if options.Sort == "Date" {
		query = append(query, fmt.Sprintf("sort=%s", options.Sort))
	}

	return strings.Join(query, "&")
}
