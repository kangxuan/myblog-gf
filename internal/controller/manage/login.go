package manage

import (
	"context"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	manage2 "myblog-gf/internal/service/manage"
)

var LoginController loginController

type loginController struct{}

func (l *loginController) Login(ctx context.Context, req *manage.LoginReq) (res *api.CommonJsonRes, err error) {
	res = manage2.LoginService.Login(ctx, req)
	return
}
