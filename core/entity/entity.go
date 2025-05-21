package entity

type Pagination[T any] struct {
	Items       []T
	TotalItems  int64
	TotalPages  int64
	CurrentPage int
	PageSize    int
}
