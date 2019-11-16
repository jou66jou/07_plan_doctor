package errors

const (
	// Success 指成功
	Success = "200001"

	// SuccessNoContent 代表成功同時不需要回傳的內容
	SuccessNoContent = "204001"

	// StatusNotModified 代表資源未變更
	StatusNotModified = "304001"

	// BadRequest 代表條件不符合
	BadRequest = "400001"

	// InvalidAuthenticationInfo 代表認證資訊錯誤
	InvalidAuthenticationInfo = "400003"

	// InvalidHeaderValue - The value provided for one of the HTTP headers was not in the correct format.")]
	InvalidHeaderValue = "400004"

	// InvalidInput - One of the request inputs is not valid.")]
	InvalidInput = "400006"

	// InvalidQueryParameterValue - An invalid value was specified for one of the query parameters in the request URI.")]
	InvalidQueryParameterValue = "400009"

	// OrderAmountNotEnabled - Order Amount Not Enabled
	OrderAmountNotEnabled = "400010"

	// MissingRequiredHeader - A required HTTP header was not specified.
	MissingRequiredHeader = "400017"

	// OutOfRangeInput - One of the request inputs is out of range.")]
	OutOfRangeInput = "400020"

	// InvalidAppVersion - Check app version from x-app-version, and the version is invalid
	InvalidAppVersion = "400030"

	// AppGeneralError - return specific error struct for APP to show
	AppGeneralError = "400040"

	// Unauthorized 指未授權
	Unauthorized = "401001"

	// AccountIsDisabled - The specified account is disabled." )]
	AccountIsDisabled = "403001"

	// AuthenticationFailed   - Server failed to authenticate the request. Make sure the value of the Authorization header is formed correctly including the signature.
	AuthenticationFailed = "403002"

	// NotAllowed - The request is understood, but it has been refused or access is not allowed.")]
	NotAllowed = "403003"

	// InsufficientAccountPermissionsWithOperation - The account being accessed does not have sufficient permissions to execute this operation
	InsufficientAccountPermissionsWithOperation = "403005"

	// UsernameOrPasswordIncorrect - Username or Password is incorrect
	UsernameOrPasswordIncorrect = "403006"

	// OtpRequired - OTP Binding is required.
	OtpRequired = "403007"

	// OtpAuthorizationRequired - Two-factor authorization is required
	OtpAuthorizationRequired = "403008"

	// OtpIncorrect - OTP is incorrect
	OtpIncorrect = "403009"

	// ResetPasswordRequired - Reset Password Required
	ResetPasswordRequired = "403010"

	// ResourceNotFound - The specified resource does not exist.
	ResourceNotFound = "404001"

	// ResourceDependencyNotFound - The specified resource dependency does not exist
	ResourceDependencyNotFound = "404002"

	// OrderNotFound - The specified order not found
	OrderNotFound = "404003"

	// SettingNotFound - The specified setting not found
	SettingNotFound = "404004"

	// QRCodeNotFound - The specified QRCode not found
	QRCodeNotFound = "404005"

	// PaymentAccountNotFound - Payment account not found
	PaymentAccountNotFound = "404006"

	// CustomerServiceNotFound - Customer service not found
	CustomerServiceNotFound = "404007"

	// OrderTypeNotFound - Order type not found
	OrderTypeNotFound = "404008"

	// QRCodeTypeNotFound - QRCode type not found
	QRCodeTypeNotFound = "404009"

	// AccountAlreadyExists - The specified account already exists.
	AccountAlreadyExists = "409001"

	// ResourceAlreadyExists - Conflict (409) - The specified resource already exists.")]
	ResourceAlreadyExists = "409004"

	// InternalError - "Internal Server Error (500) - The server encountered an internal error. Please retry the request.
	InternalError = "500001"

	// ServiceUnavailable - "Service Unavailable Error (503) - the server is currently unable to handle the request due to a temporary overload or scheduled maintenance
	ServiceUnavailable = "503001"
)
