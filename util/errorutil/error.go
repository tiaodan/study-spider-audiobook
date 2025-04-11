// 功能: 封装错误处理
package errorutil

import "log"

// 封装panic 的错误处理
/*
参数:
	err: 错误
	msg: 错误信息
返回值:
	errorCode int 错误码 1 - 正常，0 - 异常
*/
func ErrorPanic(err error, msg string) int {
	errorCode := 1
	// 异常, 有错误
	if err != nil {
		errorCode = 0
		panic(msg + ": " + err.Error())
	}
	return errorCode
}

// 封装纯打印错误
/*
参数:
	err: 错误
	msg: 错误信息
返回值:
	errorCode int 错误码 1 - 正常，0 - 异常
*/
func ErrorPrint(err error, msg string) int {
	errorCode := 1
	if err != nil {
		errorCode = 0
		log.Println(msg + ": " + err.Error())
	}
	return errorCode
}
