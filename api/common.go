package api

type CommonJsonRes struct {
	Code    int         `json:"code" dc:"状态码" d:"1"`
	Message string      `json:"message" dc:"信息" d:"message"`
	Data    interface{} `json:"data" dc:"数据"`
}
