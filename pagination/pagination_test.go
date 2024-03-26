package pagination

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaginate(t *testing.T) {
	p := Paginate(&Pagination{
		Page:         0,
		Capacity:     0,
		TotalRecords: 10,
		MinCapacity:  1,
		MaxCapacity:  100,
	})
	require.Equal(t, 1, p.GetPage())
	require.Equal(t, 1, p.GetCapacity())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 10, p.GetTotalPages())
	require.Equal(t, 0, p.GetOffset())
	require.Equal(t, 1, p.GetLimit())

	p = Paginate(&Pagination{
		Page:         1,
		Capacity:     3,
		TotalRecords: 10,
		MinCapacity:  1,
		MaxCapacity:  100,
	})
	require.Equal(t, 1, p.GetPage())
	require.Equal(t, 3, p.GetCapacity())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 4, p.GetTotalPages())
	require.Equal(t, 0, p.GetOffset())
	require.Equal(t, 3, p.GetLimit())

	p = Paginate(&Pagination{
		Page:         2,
		Capacity:     3,
		TotalRecords: 10,
		MinCapacity:  1,
		MaxCapacity:  100,
	})
	require.Equal(t, 2, p.GetPage())
	require.Equal(t, 3, p.GetCapacity())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 4, p.GetTotalPages())
	require.Equal(t, 3, p.GetOffset())
	require.Equal(t, 3, p.GetLimit())

	p = Paginate(&Pagination{
		Page:         3,
		Capacity:     3,
		TotalRecords: 10,
		MinCapacity:  1,
		MaxCapacity:  100,
	})
	require.Equal(t, 3, p.GetPage())
	require.Equal(t, 3, p.GetCapacity())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 4, p.GetTotalPages())
	require.Equal(t, 6, p.GetOffset())
	require.Equal(t, 3, p.GetLimit())

	p = Paginate(&Pagination{
		Page:         4,
		Capacity:     3,
		TotalRecords: 10,
		MinCapacity:  1,
		MaxCapacity:  100,
	})
	require.Equal(t, 4, p.GetPage())
	require.Equal(t, 3, p.GetCapacity())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 4, p.GetTotalPages())
	require.Equal(t, 9, p.GetOffset())
	require.Equal(t, 1, p.GetLimit())
}
