package public

type Page struct {
	// 页数据量
	PageSize int `json:"pageSize"`
	// 页码
	PageNum int `json:"pageNum"`
}