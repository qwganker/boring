package request

const (
	DefaultPage     = 1
	DefaultPageSize = 20
	MaxPageSize     = 100
)

// PageRequest 封装分页查询的输入参数。
type PageRequest struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

// Normalize 校正分页参数，确保在合理范围。
func (p *PageRequest) Normalize() {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultPageSize
	}
	if p.PageSize > MaxPageSize {
		p.PageSize = MaxPageSize
	}
}

// Offset 返回数据库查询的 offset。
func (p PageRequest) Offset() int {
	if p.Page <= 1 {
		return 0
	}
	return (p.Page - 1) * p.PageSize
}

// Limit 返回数据库查询的 limit。
func (p PageRequest) Limit() int {
	if p.PageSize <= 0 {
		return DefaultPageSize
	}
	return p.PageSize
}

// PageResult 为分页查询的通用返回结构。
type PageResult struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}

// NewPageResult 根据分页请求和数据构造标准返回。
func NewPageResult(req PageRequest, total int64, data interface{}) PageResult {
	return PageResult{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
		Data:     data,
	}
}
