package errors

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// 自定義的 errors, 我還要定義更多 error \(ˋ皿ˊ)/  <---怪叔叔 ಠ_ಠ
var (
	// ErrBadRequest                       =  _error{Code: "400000", Message: http.StatusText(http.StatusBadRequest), Status: http.StatusBadRequest}
	ErrInvalidInput          = _error{Code: "400001", Message: "One of the request inputs is not valid.", Status: http.StatusBadRequest}
	ErrInvalidHeaderValue    = _error{Code: "400004", Message: "The value provided for one of the HTTP headers was not in the correct format.", Status: http.StatusBadRequest}
	ErrMissingRequiredHeader = _error{Code: "400017", Message: "A required HTTP header was not specified.", Status: http.StatusBadRequest}
	ErrInternalDataNotSync   = _error{Code: "400041", Message: "The specified data not sync in system, please wait a moment.", Status: http.StatusBadRequest}

	ErrUnauthorized                = _error{Code: "401001", Message: http.StatusText(http.StatusUnauthorized), Status: http.StatusUnauthorized}
	ErrInvalidAuthenticationInfo   = _error{Code: "401001", Message: "The authentication information was not provided in the correct format. Verify the value of Authorization header.", Status: http.StatusUnauthorized}
	ErrUsernameOrPasswordIncorrect = _error{Code: "401002", Message: "Username or Password is incorrect.", Status: http.StatusUnauthorized}

	// ErrForbidden                                   =  _error{Code: "403000", Message: http.StatusText(http.StatusForbidden), Status: http.StatusForbidden}
	ErrAccountIsDisabled                           = _error{Code: "403001", Message: "The specified account is disabled.", Status: http.StatusForbidden}
	ErrAuthenticationFailed                        = _error{Code: "403002", Message: "Server failed to authenticate the request. Make sure the value of the Authorization header is formed correctly including the signature.", Status: http.StatusForbidden}
	ErrNotAllowed                                  = _error{Code: "403003", Message: "The request is understood, but it has been refused or access is not allowed.", Status: http.StatusForbidden}
	ErrOtpExpired                                  = _error{Code: "403004", Message: "OTP is expired.", Status: http.StatusForbidden}
	ErrInsufficientAccountPermissionsWithOperation = _error{Code: "403005", Message: "The account being accessed does not have sufficient permissions to execute this operation.", Status: http.StatusForbidden}
	ErrOptRequired                                 = _error{Code: "403007", Message: "OTP Binding is required.", Status: http.StatusForbidden}
	ErrOtpAuthorizationRequired                    = _error{Code: "403008", Message: "Two-factor authorization is required.", Status: http.StatusForbidden}
	ErrOtpIncorrect                                = _error{Code: "403009", Message: "OTP is incorrect.", Status: http.StatusForbidden}
	ErrResetPasswordRequired                       = _error{Code: "403010", Message: "Reset password is required.", Status: http.StatusForbidden}

	// ErrNotFound         =  _error{Code: "404000", Message: http.StatusText(http.StatusNotFound), Status: http.StatusNotFound}
	ErrResourceNotFound = _error{Code: "404001", Message: "The specified resource does not exist.", Status: http.StatusNotFound}

	ErrConflict              = _error{Code: "409000", Message: http.StatusText(http.StatusConflict), Status: http.StatusConflict}
	ErrAccountAlreadyExists  = _error{Code: "409001", Message: "The specified account already exists.", Status: http.StatusConflict}
	ErrAccountBeingCreated   = _error{Code: "409002", Message: "The specified account is in the process of being created.", Status: http.StatusConflict}
	ErrResourceAlreadyExists = _error{Code: "409004", Message: "The specified resource already exists.", Status: http.StatusConflict}

	ErrInternalServerError = _error{Code: "500000", Message: http.StatusText(http.StatusInternalServerError), Status: http.StatusInternalServerError}
	ErrInternalError       = _error{Code: "500001", Message: "The server encountered an internal error. Please retry the request.", Status: http.StatusInternalServerError}
)

type _error struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e _error) Error() string {
	var b strings.Builder
	_, _ = b.WriteRune('[')
	_, _ = b.WriteString(e.Code)
	_, _ = b.WriteRune(']')
	_, _ = b.WriteRune(' ')
	_, _ = b.WriteString(e.Message)
	return b.String()
}

// NewWithMessage 抽換錯誤訊息
// 未定義的錯誤會被視為 ErrInternalError 類型
func NewWithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	causeErr := errors.Cause(err)
	_err, ok := causeErr.(_error)
	if !ok {
		return WithStack(_error{
			Status:  ErrInternalError.Status,
			Code:    ErrInternalError.Code,
			Message: message,
		})
	}
	return WithStack(_error{
		Status:  _err.Status,
		Code:    _err.Code,
		Message: message,
	})
}
