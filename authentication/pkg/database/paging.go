package database

import "fmt"

func PagingBy(page, limit int, sort, by string) string {

	var orderBy string
	if sort == "" && sort != "desc" || sort != "asc" {
		sort = "desc"
	}

	if by == "" {
		orderBy = "created_at " + sort
	} else {
		orderBy = by + " " + sort
	}

	if page > 0 && limit > 0 {
		return fmt.Sprintf("LIMIT %d OFFSET %d %s", limit, (page-1)*limit, orderBy)
	} else {
		return orderBy
	}

}
