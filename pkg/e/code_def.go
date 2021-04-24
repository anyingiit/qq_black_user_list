package e

import "net/http"

var (
	OK              = &Error{Code: 0, DevInfo: "成功", Status: http.StatusOK, Message: "成功", MoreInfo: ""}
	Faild           = &Error{Code: -1, DevInfo: "失败", Status: http.StatusInternalServerError, Message: "失败", MoreInfo: ""}
	ErrInput        = &Error{Code: -2, DevInfo: "解析数据失败", Status: http.StatusInternalServerError, Message: "输入数据格式有误", MoreInfo: ""}
	ErrDB           = &Error{Code: -3, DevInfo: "数据库请求错误", Status: http.StatusInternalServerError, Message: "服务器内部错误", MoreInfo: ""}
	ErrSysFuncFaild = &Error{Code: -4, DevInfo: "内部函数处理错误", Status: http.StatusInternalServerError, Message: "服务器内部错误", MoreInfo: ""}
	ErrToken        = &Error{Code: -5, DevInfo: "Token错误或Token过期", Status: http.StatusUnauthorized, Message: "权限验证失败", MoreInfo: ""}

	//public response
	ErrNoUser                       = &Error{Code: 10001, DevInfo: "未找到用户", Status: http.StatusOK, Message: "用户名或密码错误", MoreInfo: ""}
	ErrPass                         = &Error{Code: 10002, DevInfo: "用户密码错误", Status: http.StatusOK, Message: "用户名或密码错误", MoreInfo: ""}
	ErrDupUser                      = &Error{Code: 10003, DevInfo: "用户已存在", Status: http.StatusOK, Message: "用户已存在", MoreInfo: ""}
	ErrRequestVerificationCodeFaild = &Error{Code: 10004, DevInfo: "验证码请求失败", Status: http.StatusOK, Message: "验证码请求失败", MoreInfo: ""}
	ErrVerificationCodeErr          = &Error{Code: 10005, DevInfo: "验证码错误", Status: http.StatusOK, Message: "验证码错误", MoreInfo: ""}
	ErrPassNotStrengthOrFormErr     = &Error{Code: 10006, DevInfo: "密码强度不符合要求", Status: http.StatusOK, Message: "密码强度不符合要求", MoreInfo: ""}

	ErrNoQqInList = &Error{Code: 20001, DevInfo: "没有查询到QQ用户", Status: http.StatusOK, Message: "没有找到该QQ用户", MoreInfo: ""}
)
