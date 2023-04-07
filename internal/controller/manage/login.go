package manage

import (
	"context"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	manage2 "myblog-gf/internal/service/manage"
)

var LoginController loginController

type loginController struct{}

// Login 登入
func (l *loginController) Login(ctx context.Context, req *manage.LoginReq) (res *api.CommonJsonRes, err error) {
	res, err = manage2.LoginService.Login(ctx, req)
	return
}

// Logout 登出
func (l *loginController) Logout(ctx context.Context, req *manage.LogoutReq) (res *api.CommonJsonRes, err error) {
	res = manage2.LoginService.Logout(ctx, req)
	return
}
