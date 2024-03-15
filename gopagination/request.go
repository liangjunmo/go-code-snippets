package gopagination

type Request interface {
	Paginate(totalRecords int) Pagination
}

type DefaultRequest struct {
	Page     int `form:"page" json:"page"`
	Capacity int `form:"capacity" json:"capacity"`
}

func (req *DefaultRequest) Paginate(totalRecords int) Pagination {
	return Paginate(req.Page, req.Capacity, totalRecords, 1, 10)
}
