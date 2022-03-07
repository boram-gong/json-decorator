package body

var (
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

func NewCommonResp() *CommonResp {
	respBody := new(CommonResp)
	respBody.RespHead = SuccessResp
	return respBody
}
