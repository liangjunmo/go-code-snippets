package pagination

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaginate(t *testing.T) {
	p := Paginate(0, 0, 10, 1, 10)
	require.Equal(t, 1, p.GetPage())
	require.Equal(t, 10, p.GetCapacityPerPage())
	require.Equal(t, 1, p.GetTotalPages())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 0, p.GetOffset())
	require.Equal(t, 10, p.GetLimit())

	p = Paginate(1, 3, 10, 1, 10)
	require.Equal(t, 1, p.GetPage())
	require.Equal(t, 3, p.GetCapacityPerPage())
	require.Equal(t, 4, p.GetTotalPages())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 0, p.GetOffset())
	require.Equal(t, 3, p.GetLimit())

	p = Paginate(2, 3, 10, 1, 10)
	require.Equal(t, 2, p.GetPage())
	require.Equal(t, 3, p.GetCapacityPerPage())
	require.Equal(t, 4, p.GetTotalPages())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 3, p.GetOffset())
	require.Equal(t, 3, p.GetLimit())

	p = Paginate(3, 3, 10, 1, 10)
	require.Equal(t, 3, p.GetPage())
	require.Equal(t, 3, p.GetCapacityPerPage())
	require.Equal(t, 4, p.GetTotalPages())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 6, p.GetOffset())
	require.Equal(t, 3, p.GetLimit())

	p = Paginate(4, 3, 10, 1, 10)
	require.Equal(t, 4, p.GetPage())
	require.Equal(t, 3, p.GetCapacityPerPage())
	require.Equal(t, 4, p.GetTotalPages())
	require.Equal(t, 10, p.GetTotalRecords())
	require.Equal(t, 9, p.GetOffset())
	require.Equal(t, 1, p.GetLimit())
}
