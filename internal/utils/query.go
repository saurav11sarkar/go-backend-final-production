package utils

import (
	"net/url"
	"strconv"
	"strings"
)

type Query struct {
	Page       int
	Limit      int
	Search     string
	SortBy     string
	SortOrder  string
	Role       string
	Status     string
	CategoryID string
}

func ParseQuery(v url.Values) Query {
	q := Query{Page: 1, Limit: 10, Search: strings.TrimSpace(v.Get("search")), SortBy: "created_at", SortOrder: "desc", Role: v.Get("role"), Status: v.Get("status"), CategoryID: v.Get("categoryId")}
	if n, err := strconv.Atoi(v.Get("page")); err == nil && n > 0 {
		q.Page = n
	}
	if n, err := strconv.Atoi(v.Get("limit")); err == nil && n > 0 && n <= 100 {
		q.Limit = n
	}
	if s := v.Get("sortBy"); s != "" {
		q.SortBy = s
	}
	if s := strings.ToLower(v.Get("sortOrder")); s == "asc" || s == "desc" {
		q.SortOrder = s
	}
	return q
}
func (q Query) Offset() int { return (q.Page - 1) * q.Limit }
