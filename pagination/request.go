package pagination

type Request interface {
	Paginate(totalRecords int) Result
}

type DefaultRequest struct {
	Page     int `form:"page" json:"page"`
	Capacity int `form:"capacity" json:"capacity"`
}

func (req *DefaultRequest) Paginate(totalRecords int) Result {
	return Paginate(&Pagination{
		Page:         req.Page,
		Capacity:     req.Capacity,
		TotalRecords: totalRecords,
		MinCapacity:  1,
		MaxCapacity:  100,
	})
}
