package pagination

func Paginate(page, capacity, totalRecords, defaultPage, defaultCapacity int) Pagination {
	if page < 1 {
		page = defaultPage
	}
	if capacity < 1 {
		capacity = defaultCapacity
	}

	totalPages := totalRecords / capacity
	if totalRecords%capacity > 0 {
		totalPages++
	}
	if totalPages == 0 {
		totalPages = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * capacity
	limit := capacity
	if offset+limit > totalRecords {
		limit = totalRecords - offset
	}
	return &DefaultPagination{
		Page:            page,
		CapacityPerPage: capacity,
		TotalPages:      totalPages,
		TotalRecords:    totalRecords,
		Offset:          offset,
		Limit:           limit,
	}
}
