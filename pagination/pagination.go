package pagination

import "github.com/spf13/cast"

const (
	defaultPageIndex int = 1
	defaultPageSize  int = 10
	maxPageSize      int = 500
)

type Pagination struct {
	PageIndex    int `json:"page_index"`
	PageSize     int `json:"page_size"`
	TotalPages   int `json:"page_total"`
	TotalResults int `json:"result_total"`
	Offset       int `json:"-"`
	Limit        int `json:"-"`
}

type Request struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

func (req Request) Paginate(count int64) Pagination {
	if req.PageIndex < 1 {
		req.PageIndex = defaultPageIndex
	}
	if req.PageSize < 1 {
		req.PageSize = defaultPageSize
	}
	if req.PageSize > maxPageSize {
		req.PageSize = maxPageSize
	}

	var (
		totalResults = cast.ToInt(count)
		totalPages   = totalResults / req.PageSize
		offset       int
		limit        int
	)
	if totalResults%req.PageSize > 0 {
		totalPages += 1
	}
	if totalPages == 0 {
		totalPages = 1
	}
	if req.PageIndex > totalPages {
		req.PageIndex = totalPages
	}

	offset = (req.PageIndex - 1) * req.PageSize
	limit = req.PageSize
	if offset+limit > totalResults {
		limit = totalResults - offset
	}

	return Pagination{
		PageIndex:    req.PageIndex,
		PageSize:     req.PageSize,
		TotalPages:   totalPages,
		TotalResults: totalResults,
		Offset:       offset,
		Limit:        limit,
	}
}
