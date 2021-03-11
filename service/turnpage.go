package service

type Pagination struct {
	Limit  int
	Offset int
}
type Page struct {
	PageSize   int
	PageNumber int
}

func (p Page) Limit() int {
	return p.PageSize
}
func (p Page) Offset() int {
	return (p.PageNumber - 1) * (p.PageSize)
}
