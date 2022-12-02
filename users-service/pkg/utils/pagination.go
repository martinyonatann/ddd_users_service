package utils

import "strconv"

const (
	defaultSize = 10
)

type PaginationQuery struct {
	Size    int    `json:"size,omitempty"`
	Page    int    `json:"page,omitempty"`
	OrderBy string `json:"orderBy,omitempty"`
}

// Set page size
func (q *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultSize
		return nil
	}

	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}

	q.Size = n

	return nil
}
