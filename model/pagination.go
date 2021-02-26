package model

// Pagination 分页
type Pagination struct {
	Page   int         `json:"page" query:"page"`
	Size   int         `json:"size" query:"size"`
	Offset int         `json:"-"`
	Total  int64       `json:"total"`
	Data   interface{} `json:"data"`
}

// NewPagination 创建分页对象
func NewPagination(page int, size int) *Pagination {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	offset := (page - 1) * size
	return &Pagination{Page: page, Size: size, Offset: offset}
}

// Set 设置数据
func (pagination *Pagination) Set(total int64, data interface{}) {
	pagination.Total = total
	pagination.Data = data
}
