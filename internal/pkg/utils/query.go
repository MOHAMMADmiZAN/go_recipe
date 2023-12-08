package utils

import (
	"fmt"
	"net/url"
	"strconv"
)

// Config represents the default configuration values.
type Config struct {
	TotalItems int
	Limit      int
	Page       int
	SortType   string
	Sort       string
	Search     string
}

// DefaultConfig contains the default configuration values.
var DefaultConfig = Config{
	TotalItems: 0,
	Limit:      10,
	Page:       1,
	SortType:   "dsc",
	Sort:       "updatedAt",
	Search:     "",
}

// PaginationOptions struct represents pagination options.
type PaginationOptions struct {
	TotalItems int
	Limit      int
	Page       int
}

// GetPagination calculates pagination information.
func GetPagination(options PaginationOptions) map[string]interface{} {
	page := options.Page
	limit := options.Limit
	totalItems := options.TotalItems
	totalPage := (totalItems + limit - 1) / limit
	hasNext := page < totalPage
	hasPrev := page > 1

	pagination := map[string]interface{}{
		"page":       page,
		"limit":      limit,
		"totalItems": totalItems,
		"totalPage":  totalPage,
	}

	if hasNext {
		pagination["next"] = page + 1
	}
	if hasPrev {
		pagination["prev"] = page - 1
	}

	return pagination
}

// HATEOASOptions struct represents options for generating HATEOAS links.
type HATEOASOptions struct {
	URL     string
	Path    string
	Query   map[string]string
	Page    int
	HasNext bool
	HasPrev bool
}

// GetHATEOASForAllItems generates HATEOAS links for pagination.
func GetHATEOASForAllItems(options HATEOASOptions) map[string]string {
	query := options.Query
	query["page"] = strconv.Itoa(options.Page)

	queryStr := GenerateQueryString(query)
	selfLink := fmt.Sprintf("%s%s?%s", options.URL, options.Path, queryStr)

	links := map[string]string{
		"self": selfLink,
	}

	if options.HasNext {
		query["page"] = strconv.Itoa(options.Page + 1)
		nextQueryStr := GenerateQueryString(query)
		nextLink := fmt.Sprintf("%s%s?%s", options.URL, options.Path, nextQueryStr)
		links["next"] = nextLink
	}

	if options.HasPrev {
		query["page"] = strconv.Itoa(options.Page - 1)
		prevQueryStr := GenerateQueryString(query)
		prevLink := fmt.Sprintf("%s%s?%s", options.URL, options.Path, prevQueryStr)
		links["prev"] = prevLink
	}

	return links
}

// QueryParams struct represents parsed and validated query parameters.
type QueryParams struct {
	Skip   int
	Limit  int
	Sort   map[string]int
	Search string
	Page   int
}

// ParseQueryParams parses and validates query parameters for pagination, sorting, and search.
func ParseQueryParams(queryParams url.Values) QueryParams {
	pageStr := queryParams.Get("page")
	limitStr := queryParams.Get("limit")
	sortField := queryParams.Get("sort")
	sortType := queryParams.Get("sort_type")
	search := queryParams.Get("search")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = DefaultConfig.Limit
	}

	sort := map[string]int{}
	if sortField != "" {
		sortTypeInt := 1
		if sortType == "dsc" {
			sortTypeInt = -1
		}
		sort[sortField] = sortTypeInt
	}

	return QueryParams{
		Skip:   (page - 1) * limit,
		Limit:  limit,
		Sort:   sort,
		Search: search,
		Page:   page,
	}
}
