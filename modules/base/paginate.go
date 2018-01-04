package base

const PerPage = 2

type Pagination struct {
	TotalCount int
	CurPage    int
	Page       float64
	PageCount  int
}
