package utils

import (
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/beego/beego/v2/server/web/context"
)

type Paginator struct {
	Ctx         *context.Context
	Request     *http.Request
	PerPageNums int
	MaxPages    int

	nums      int64
	pageRange []int
	pageNums  int
}

func (p *Paginator) PageNums() int {
	if p.pageNums != 0 {
		return p.pageNums
	}
	pageNums := math.Ceil(float64(p.nums) / float64(p.PerPageNums))
	if p.MaxPages > 0 {
		pageNums = math.Min(pageNums, float64(p.MaxPages))
	}
	p.pageNums = int(pageNums)
	return p.pageNums
}

func (p *Paginator) Nums() int64 {
	return p.nums
}

func (p *Paginator) SetNums(nums interface{}) {
	p.nums, _ = ToInt64(nums)
}

func (p *Paginator) Page() int {
	page, _ := strconv.Atoi(p.Ctx.Input.Param(":page"))

	if page > p.PageNums() {
		page = p.PageNums()
	}
	if page <= 0 {
		page = 1
	}
	return page
}

func (p *Paginator) Pages() []int {
	if p.pageRange != nil || p.nums <= 0 {
		return p.pageRange
	}

	pageNums := p.PageNums()
	currentPage := p.Page()

	// Define the edges based on the current page.
	startEdge, endEdge := 1, pageNums
	edgeDisplay := 2

	// Calculate the window of pages around the current page.
	startPage := max(startEdge, currentPage-2)
	endPage := min(currentPage+2, pageNums)

	pages := make([]int, 0, endEdge)

	// Always include the start edge pages.
	for i := startEdge; i <= min(edgeDisplay, pageNums); i++ {
		pages = append(pages, i)
	}

	// Use ellipsis to indicate skipped pages between edges and middle pages.
	if startPage > edgeDisplay+1 {
		pages = append(pages, 0)
	}

	// Add the window of pages around the current page.
	for i := max(startPage, edgeDisplay+1); i <= min(endPage, pageNums-edgeDisplay); i++ {
		pages = append(pages, i)
	}

	// Use ellipsis to indicate skipped pages between middle pages and end edge.
	if endPage < pageNums-edgeDisplay {
		pages = append(pages, 0)
	}

	// Always include the end edge pages.
	for i := max(pageNums-edgeDisplay+1, endPage+1); i <= endEdge; i++ {
		pages = append(pages, i)
	}

	p.pageRange = pages
	return p.pageRange
}

func (p *Paginator) PageLink(page int) string {
	link, _ := url.ParseRequestURI(p.Request.RequestURI)
	values := link.Query()
	if page == 1 {
		values.Del("p")
	} else {
		values.Set("p", strconv.Itoa(page))
	}
	link.RawQuery = values.Encode()
	return strconv.Itoa(page)
}

func (p *Paginator) PageLinkPrev() (link string) {
	if p.HasPrev() {
		link = p.PageLink(p.Page() - 1)
	}
	return
}

func (p *Paginator) PageLinkNext() (link string) {
	if p.HasNext() {
		link = p.PageLink(p.Page() + 1)
	}
	return
}

func (p *Paginator) HasPrev() bool {
	return p.Page() > 1
}

func (p *Paginator) HasNext() bool {
	return p.Page() < p.PageNums()
}

func (p *Paginator) IsActive(page int) bool {
	return p.Page() == page
}

// func (p *Paginator) Offset() int {
// 	return (p.Page() - 1) * p.PerPageNums
// }

func (p *Paginator) HasPages() bool {
	return p.PageNums() > 1
}

func NewPaginator(ctx *context.Context, per int, nums interface{}) *Paginator {
	p := Paginator{}
	p.Request = ctx.Request
	p.Ctx = ctx
	if per <= 0 {
		per = 10
	}
	p.PerPageNums = per
	p.SetNums(nums)
	return &p
}
