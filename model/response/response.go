package response

import (
	"my-backend/model/code"
)

type ResponseData struct {
	Code code.ResCode `json:"code"`
	Msg  interface{}  `json:"msg"`
	Data interface{}  `json:"data,omitempty"` //没有值忽略
}

func ResponseError(code code.ResCode) ResponseData {
	return ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
}

func ResponseErrorWithMsg(code code.ResCode, msg interface{}) ResponseData {
	return ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func ResponseSuccess(data interface{}) ResponseData {
	return ResponseData{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}
