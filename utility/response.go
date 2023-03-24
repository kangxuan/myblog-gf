package utility

import "myblog-gf/api"

var CommonResponse = commonResponse{}

type commonResponse struct {
}

func (c *commonResponse) ErrorMsg(msg string) (res *api.CommonJsonRes) {
	return &api.CommonJsonRes{
		Code:    0,
		Message: msg,
		Data:    nil,
	}
}

func (c *commonResponse) SuccessMsg(msg string, data interface{}) (res *api.CommonJsonRes) {
	return &api.CommonJsonRes{
		Code:    1,
		Message: msg,
		Data:    data,
	}
}
