package manage

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"myblog-gf/api"
	"myblog-gf/api/v1/manage"
	"myblog-gf/utility"
)

var LoginService loginService

type loginService struct {
}

func (l *loginService) Login(ctx context.Context, req *manage.LoginReq) (res *api.CommonJsonRes) {
	rec, err := g.Model("user").Where("account = ?", req.Account).One()
	if err != nil {
		panic(err)
	}
	if rec.IsEmpty() {
		return utility.CommonResponse.ErrorMsg("账号不存在")
	}
	user := rec.Map()

	md5String, err := gmd5.Encrypt(req.Password + gconv.String(user["salt"]))
	if err != nil {
		panic(err)
	}
	if md5String != user["password"] {
		return utility.CommonResponse.ErrorMsg("密码错误")
	}

	// 设置token

	return utility.CommonResponse.SuccessMsg("登入成功", nil)
}
