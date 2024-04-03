package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginate(t *testing.T) {
	var (
		req Request
		p   Pagination
	)

	req = Request{
		PageIndex: 0,
		PageSize:  0,
	}
	p = req.Paginate(10)
	assert.Equal(t, defaultPageIndex, p.PageIndex)
	assert.Equal(t, defaultPageSize, p.PageSize)

	req = Request{
		PageIndex: 1,
		PageSize:  3,
	}
	p = req.Paginate(10)
	assert.Equal(t, 1, p.PageIndex)
	assert.Equal(t, 3, p.PageSize)
	assert.Equal(t, 4, p.TotalPages)
	assert.Equal(t, 10, p.TotalResults)
	assert.Equal(t, 0, p.Offset)
	assert.Equal(t, 3, p.Limit)

	req = Request{
		PageIndex: 2,
		PageSize:  3,
	}
	p = req.Paginate(10)
	assert.Equal(t, 2, p.PageIndex)
	assert.Equal(t, 3, p.PageSize)
	assert.Equal(t, 4, p.TotalPages)
	assert.Equal(t, 10, p.TotalResults)
	assert.Equal(t, 3, p.Offset)
	assert.Equal(t, 3, p.Limit)

	req = Request{
		PageIndex: 3,
		PageSize:  3,
	}
	p = req.Paginate(10)
	assert.Equal(t, 3, p.PageIndex)
	assert.Equal(t, 3, p.PageSize)
	assert.Equal(t, 4, p.TotalPages)
	assert.Equal(t, 10, p.TotalResults)
	assert.Equal(t, 6, p.Offset)
	assert.Equal(t, 3, p.Limit)

	req = Request{
		PageIndex: 4,
		PageSize:  3,
	}
	p = req.Paginate(10)
	assert.Equal(t, 4, p.PageIndex)
	assert.Equal(t, 3, p.PageSize)
	assert.Equal(t, 4, p.TotalPages)
	assert.Equal(t, 10, p.TotalResults)
	assert.Equal(t, 9, p.Offset)
	assert.Equal(t, 1, p.Limit)
}
