package util

import "apiserver/pkg/constvar"

type PageSetting struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
	Page   uint64 `json:"page"`
}
type Page struct {
	CurrentPage uint64      `json:"current_page"`
	PageSize    uint64      `json:"page_size"`
	TotalPage   uint64      `json:"total_page"`
	TotalCount  uint64      `json:"total_count"`
	FirstPage   bool        `json:"first_page"`
	LastPage    bool        `json:"last_page"`
	List        interface{} `json:"list"`
}

func (s *PageSetting) Setting(page uint64) {
	if s.Limit == 0 {
		s.Limit = constvar.DefaultLimit
	}
	if page == 0 {
		page = 1
	}
	s.Page = page
	s.Offset = (page - 1) * s.Limit
}

func PageUtil(count uint64, currentPage uint64, pageSize uint64, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{CurrentPage: currentPage, PageSize: pageSize, TotalPage: tp, TotalCount: count,
		FirstPage: currentPage == 1, LastPage: currentPage == tp, List: list}
}
