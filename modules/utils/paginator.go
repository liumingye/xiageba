package utils

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/server/web/context"
)

type Paginator struct {
	Ctx         *context.Context
	Request     *http.Request
	PerPageNums int
	MaxPages    int
	LinkFormat  string

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

	page = int(math.Max(1, math.Min(float64(p.PageNums()), float64(page))))

	return page
}

func (p *Paginator) Pages() []int {
	if p.pageRange != nil {
		return p.pageRange
	}

	pageNums := p.PageNums()
	currentPage := p.Page()

	pages := make([]int, 0, pageNums)

	// Add the first two pages.
	pages = append(pages, 1, 2)

	// Add the ellipsis if needed.
	if currentPage > 5 {
		pages = append(pages, 0)
	}

	// Add the window of pages around the current page.
	for i := max(3, currentPage-3); i <= min(pageNums-2, currentPage+5); i++ {
		pages = append(pages, i)
	}

	// Add the ellipsis if needed.
	if currentPage < pageNums-6 {
		pages = append(pages, 0)
	}

	// Add the last two pages.
	pages = append(pages, pageNums-1, pageNums)

	p.pageRange = pages
	return p.pageRange
}

func (p *Paginator) PageLink(page int) string {
	return fmt.Sprintf(p.LinkFormat, (strconv.Itoa(page)))
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

func (p *Paginator) HasPages() bool {
	return p.PageNums() > 1
}

func NewPaginator(ctx *context.Context, per int, nums interface{}, linkFormat string) *Paginator {
	p := Paginator{}
	p.Ctx = ctx
	p.Request = ctx.Request
	p.LinkFormat = linkFormat
	if per <= 0 {
		per = 10
	}
	p.PerPageNums = per
	p.SetNums(nums)
	return &p
}
