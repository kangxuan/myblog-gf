package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"myblog-gf/internal/consts"
	"net/http"
)

// CrossDomain 跨域请求中间件
func CrossDomain(r *ghttp.Request) {
	r.Response.CORSDefault()

	// 前置中间件
	r.Middleware.Next()
}

// ManageAuth 后台登录检验中间件
func ManageAuth(r *ghttp.Request) {
	cookie := r.Cookie.Get(consts.ManageToken)
	if cookie.IsEmpty() {
		panic("您还未登录")
	}
	r.Middleware.Next()
}

// JsonResponse 统一返回Json格式中间件
func JsonResponse(r *ghttp.Request) {
	// 后置中间件
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	//var res *api.CommonJsonRes

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}

	r.Response.WriteJson(res)
}
