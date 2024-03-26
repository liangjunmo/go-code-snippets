package pagination

type Pagination struct {
	Page         int
	Capacity     int
	TotalRecords int
	MinCapacity  int
	MaxCapacity  int
}

func Paginate(page *Pagination) Result {
	if page.Page < 1 {
		page.Page = 1
	}
	if page.Capacity < page.MinCapacity {
		page.Capacity = page.MinCapacity
	}
	if page.Capacity > page.MaxCapacity {
		page.Capacity = page.MaxCapacity
	}

	totalPages := page.TotalRecords / page.Capacity
	if page.TotalRecords%page.Capacity > 0 {
		totalPages++
	}
	if totalPages == 0 {
		totalPages = 1
	}
	if page.Page > totalPages {
		page.Page = totalPages
	}

	offset := (page.Page - 1) * page.Capacity
	limit := page.Capacity
	if offset+limit > page.TotalRecords {
		limit = page.TotalRecords - offset
	}
	return &DefaultResult{
		Page:         page.Page,
		Capacity:     page.Capacity,
		TotalRecords: page.TotalRecords,
		TotalPages:   totalPages,
		Offset:       offset,
		Limit:        limit,
	}
}

type Result interface {
	GetPage() int
	GetCapacity() int
	GetTotalRecords() int
	GetTotalPages() int
	GetOffset() int
	GetLimit() int
}

type DefaultResult struct {
	Page         int `json:"page"`
	Capacity     int `json:"capacity"`
	TotalRecords int `json:"total_records"`
	TotalPages   int `json:"total_pages"`
	Offset       int `json:"offset"`
	Limit        int `json:"limit"`
}

func (r *DefaultResult) GetPage() int {
	return r.Page
}

func (r *DefaultResult) GetCapacity() int {
	return r.Capacity
}

func (r *DefaultResult) GetTotalRecords() int {
	return r.TotalRecords
}

func (r *DefaultResult) GetTotalPages() int {
	return r.TotalPages
}

func (r *DefaultResult) GetOffset() int {
	return r.Offset
}

func (r *DefaultResult) GetLimit() int {
	return r.Limit
}
