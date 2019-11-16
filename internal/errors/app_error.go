package errors

import (
	"fmt"
)

// TODO: 整理errors 包，以error.go為主，理論上要移除apierr、app_error、error_code

// AppError 用來處理 identity app 的業務邏輯錯誤
type AppError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s-%s", e.Code, e.Message)
}

// NewAppErr functions create a new appError instance (他會被棄用)
func NewAppErr(code, message string, a ...interface{}) AppError {
	if len(a) == 0 {
		return AppError{Code: code, Message: message}
	}
	return AppError{Code: code, Message: fmt.Sprintf(message, a...)}
}
