package gopagination

type Pagination interface {
	GetPage() int
	GetCapacityPerPage() int
	GetTotalPages() int
	GetTotalRecords() int
	GetOffset() int
	GetLimit() int
}

type DefaultPagination struct {
	Page            int `json:"page"`
	CapacityPerPage int `json:"capacity_per_page"`
	TotalPages      int `json:"total_pages"`
	TotalRecords    int `json:"total_records"`
	Offset          int `json:"offset"`
	Limit           int `json:"limit"`
}

func (p *DefaultPagination) GetPage() int {
	return p.Page
}

func (p *DefaultPagination) GetCapacityPerPage() int {
	return p.CapacityPerPage
}

func (p *DefaultPagination) GetTotalPages() int {
	return p.TotalPages
}

func (p *DefaultPagination) GetTotalRecords() int {
	return p.TotalRecords
}

func (p *DefaultPagination) GetOffset() int {
	return p.Offset
}

func (p *DefaultPagination) GetLimit() int {
	return p.Limit
}
