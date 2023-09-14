package pkg

import "math"

type Paginator[T any] struct {
	Items []T `json:"items"`
	Total int `json:"total"`
	PaginatorInfo
}

type PaginatorInfo struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	Limit     int `json:"limit"`
}

type Filters[T any] struct {
	PaginatorInfo
	Search T
}

type Pagination struct {
	Limit      int   `json:"limit,omitempty;query:limit"`
	Page       int   `json:"page,omitempty;query:page"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
	Items      any   `json:"items,omitempty"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetTotalPages() int {
	return int(math.Ceil(float64(p.TotalItems) / float64(p.Limit)))
}

func (p *Pagination) Mount(model any) *Pagination {
	p.TotalPages = p.GetTotalPages()
	p.Items = model
	return p
}
