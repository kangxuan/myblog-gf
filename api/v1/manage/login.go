package manage

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"POST" tags:"登录" summary:"登入"`
	Account  string `json:"account" v:"required|max-length:10#账号必传|账号长度最长为10" dc:"账号"`
	Password string `json:"password" v:"required|max-length:50#密码必传|密码不能超过50" dc:"密码"`
}

type LogoutReq struct {
	g.Meta `path:"/login" method:"DELETE" tags:"登录" summary:"登出"`
}
