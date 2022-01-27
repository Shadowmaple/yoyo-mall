package errno

var (
	// Common errors
	OK = &Errno{Code: 0, Message: "OK"}

	InternalError     = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind           = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrValidation     = &Errno{Code: 10003, Message: "Validation failed."}
	ErrDatabase       = &Errno{Code: 10004, Message: "Database error."}
	ErrGetQuery       = &Errno{Code: 10005, Message: "Error occurred while getting query. "}
	ErrGetParam       = &Errno{Code: 10006, Message: "Error occurred while getting path params. "}
	ErrJsonUnmarshal  = &Errno{Code: 10007, Message: "Error occurred while unmarshaling json."}
	ErrRecordNotFound = &Errno{Code: 10008, Message: "Record not found in db."}

	// Auth errors
	ErrAuthFailed   = &Errno{Code: 20001, Message: "The sid or password is incorrect."}
	ErrTokenInvalid = &Errno{Code: 20002, Message: "The token is invalid."}

	// User errors
	ErrUserNotFound = &Errno{Code: 21001, Message: "User not found."}
	ErrWechatServer = &Errno{Code: 21002, Message: "Wechat server error."}

	// product errors

	// comment errors

	// order errors

	// message errors

	// logistic errors

	// cart errors

	// collection errors

	// address errors

	// feedback errors

)