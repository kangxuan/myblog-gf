package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"myblog-gf/api"
	"myblog-gf/internal/consts"
	"myblog-gf/utility"
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

	var (
		msg  string
		err  = r.GetError()
		res  *api.CommonJsonRes
		code = gerror.Code(err)
	)

	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		//测试环境
		msg = err.Error()
		//正式环境
		//msg = "服务器居然开小差了，请稍后再试吧！"
		//记录日志
		g.Log("exception").Error(gctx.New(), err)
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
	if code.Code() != 0 {
		res = utility.CommonResponse.ErrorMsg(msg)
	}
	r.Response.WriteJson(res)
}
