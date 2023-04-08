package manage

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	"myblog-gf/internal/consts"
	"myblog-gf/utility"
)

var LoginService loginService

type loginService struct{}

// Login 登录
func (l *loginService) Login(ctx context.Context, req *manage.LoginReq) (res *api.CommonJsonRes, err error) {
	rec, err := g.Model("user1").Where("account = ?", req.Account).One()
	if err != nil {
		return
	}
	if rec.IsEmpty() {
		res = utility.CommonResponse.ErrorMsg("账号不存在")
		return
	}
	user := rec.Map()

	md5String, err := gmd5.Encrypt(req.Password + gconv.String(user["salt"]))
	if err != nil {
		return
	}
	if md5String != user["password"] {
		res = utility.CommonResponse.ErrorMsg("密码错误")
		return
	}

	// 设置token
	token, err := gmd5.Encrypt(user)
	if err != nil {
		return
	}

	// 通过cxt获取request数据
	g.RequestFromCtx(ctx).Cookie.Set(consts.ManageToken, token)
	res = utility.CommonResponse.SuccessMsg("登入成功", nil)
	return
}

// Logout 登出
func (l *loginService) Logout(ctx context.Context, _ *manage.LogoutReq) (res *api.CommonJsonRes) {
	cookie := g.RequestFromCtx(ctx).Cookie
	manageToken := cookie.Get(consts.ManageToken)
	if !manageToken.IsEmpty() {
		cookie.Remove(consts.ManageToken)
	}
	return utility.CommonResponse.SuccessMsg("登出成功", nil)
}
