package errno

var (
	// Common errors
	OK = &Errno{Code: 0, Message: "OK"}

	InternalError     = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind           = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrValidation     = &Errno{Code: 10003, Message: "Validation failed."}
	ErrDatabase       = &Errno{Code: 10004, Message: "Database error."}
	ErrGetQuery       = &Errno{Code: 10005, Message: "Error occurred while getting query."}
	ErrGetParam       = &Errno{Code: 10006, Message: "Error occurred while getting path params."}
	ErrJsonUnmarshal  = &Errno{Code: 10007, Message: "Error occurred while unmarshaling json."}
	ErrRecordNotFound = &Errno{Code: 10008, Message: "Record not found in db."}
	ErrGetFile        = &Errno{Code: 10009, Message: "Error occurred while getting file."}

	// Auth errors
	ErrAuthFailed   = &Errno{Code: 20001, Message: "The sid or password is incorrect."}
	ErrTokenInvalid = &Errno{Code: 20002, Message: "The token is invalid."}
	ErrPwdWrong     = &Errno{Code: 20003, Message: "Password is wrong."}

	// User errors
	ErrUserNotFound = &Errno{Code: 21001, Message: "User not found."}
	ErrWechatServer = &Errno{Code: 21002, Message: "Wechat server error."}

	// order errors
	ErrOrderExpectedStatus = &Errno{Code: 22001, Message: "Expected updated order'status is wrong."}

	// coupon errors
	ErrCouponCodeWrong  = &Errno{Code: 23001, Message: "Coupon's code is Wrong."}
	ErrCouponGrabbed    = &Errno{Code: 23002, Message: "Has already grabbed."}
	ErrCouponNotPublic  = &Errno{Code: 23003, Message: "The coupon is not public."}
	ErrCouponNotExist   = &Errno{Code: 23004, Message: "The coupon does not exist."}
	ErrCouponCanNotGrab = &Errno{Code: 23005, Message: "Can not grab the coupon now."}

	// product errors

	// comment errors

	// message errors

	// logistic errors

	// cart errors

	// collection errors

	// address errors

	// feedback errors

)
