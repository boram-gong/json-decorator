package body

import "sync"

var (
	RespPool    = sync.Pool{New: newCommonResp}
	SuccessResp = &RespHead{
		Code: 200,
		Msg:  "成功",
	}
)

type RespHead struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CommonResp struct {
	*RespHead
	Data interface{} `json:"data"`
}

func (r *CommonResp) FailResp(code int, err string) {
	r.RespHead = &RespHead{
		Code: code,
		Msg:  err,
	}
}

func (r *CommonResp) Init() {
	r.RespHead = SuccessResp
	r.Data = nil
}

func newCommonResp() interface{} {
	return new(CommonResp)
}
