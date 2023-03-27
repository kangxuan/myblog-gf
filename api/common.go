package api

type CommonJsonRes struct {
	Code    int         `json:"code" dc:"状态码" d:"1"`
	Message string      `json:"message" dc:"信息" d:"message"`
	Data    interface{} `json:"data" dc:"数据"`
}

type PageParams struct {
	Page     int `p:"page" dc:"页码" v:"min:1#页面数不能小于1" d:"1"`
	PageSize int `p:"page_size" dc:"每页数量" v:"min:1#页数不能小于1" d:"10"`
}
