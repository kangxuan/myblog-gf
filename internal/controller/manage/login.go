package manage

import (
	"context"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	manage2 "myblog-gf/internal/service/manage"
)

var LoginController loginController

type loginController struct {
}

func (l *loginController) login(ctx context.Context, req *manage.LoginReq) (res *api.CommonJsonRes, err error) {
	//var reqParams *manage.LoginReq
	//if err := gconv.Struct(req, &reqParams); err != nil {
	//	return nil, err
	//}
	res = manage2.LoginService.Login(ctx, req)
	return
}
