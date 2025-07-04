package pagination

import (
	"fmt"
	"math"
)

const dfPageSize = 100
const maxPageSize = math.MaxInt64
const firstPage = 1

type Page struct {
	Size int64 `json:"size"` // Page Size, default 100
	Numb int64 `json:"numb"` // Page Numb, From One
}

func PageAll() *Page {
	return &Page{
		Size: maxPageSize,
		Numb: firstPage,
	}
}

func PageNormal() *Page {
	return NewPage(dfPageSize, firstPage)
}

func NewPage(size int, numb int) *Page {
	if size <= 0 {
		size = dfPageSize
	}
	if numb <= 0 {
		numb = firstPage
	}
	return &Page{
		Size: int64(size),
		Numb: int64(numb),
	}
}

type Paging struct {
	Size  int64 `json:"size"`  // Page Size, default 100
	Numb  int64 `json:"numb"`  // Page Numb, From One
	Total int64 `json:"total"` // Page Total, The Page Total
	Count int64 `json:"count"` // Item Count, The Item Count
}

func (paging *Paging) ToString() string {
	return fmt.Sprintf("size: %d|numb: %d|total: %d|count: %d",
		paging.Size, paging.Numb, paging.Total, paging.Count)
}

func (paging *Paging) WithCount(count int64) *Paging {
	if paging.Size == 0 {
		paging.Size = dfPageSize
	}
	if paging.Numb < firstPage {
		paging.Numb = firstPage
	}

	if count == 0 {
		paging.Total = 0
		paging.Numb = firstPage
		return paging
	}
	paging.Count = count
	paging.Total = paging.Count / paging.Size
	if paging.Count%paging.Size != 0 {
		paging.Total += 1
	}

	if paging.Numb > paging.Total {
		paging.Numb = paging.Total
	}

	if paging.Numb < 1 {
		paging.Numb = 1
	}
	return paging
}

func (paging *Paging) Skip() int64 {
	return (paging.Numb - 1) * paging.Size
}

func (paging *Paging) Limit() int64 {
	return paging.Size
}

func PagingOfPage(page *Page) *Paging {
	if page == nil {
		return NewPaging(dfPageSize, firstPage)
	}
	return NewPaging(page.Size, page.Numb)
}

func NewPaging(size int64, current int64) *Paging {
	if size <= 0 {
		size = dfPageSize
	}
	if current < firstPage {
		current = firstPage
	}
	return &Paging{
		Size:  size,
		Numb:  current,
		Total: 0,
		Count: 0,
	}
}

func PagingAll() *Paging {
	return NewPaging(maxPageSize, firstPage)
}

type Pagination[T any] struct {
	Paging *Paging `json:"paging"`
	Data   []T     `json:"data"`
}

func NewPagination[T any](paging *Paging, data []T) *Pagination[T] {
	if data == nil {
		data = make([]T, 0)
	}
	return &Pagination[T]{
		Paging: paging,
		Data:   data,
	}
}

func (p *Pagination[T]) Iter(fn func(item T, idx int)) {
	for idx, item := range p.Data {
		fn(item, idx)
	}
}
