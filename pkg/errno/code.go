package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation    = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase      = &Errno{Code: 20002, Message: "Database error."}
	ErrToken         = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrCasbin        = &Errno{Code: 20004, Message: "rbac casbin error"}
	ErrNotPermission = &Errno{Code: 20005, Message: "rbac no permission"}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrUserExist         = &Errno{Code: 20105, Message: "User already exists"}
	ErrNotUserExist      = &Errno{Code: 20106, Message: "User does not exist"}
	ErrUsernameExist     = &Errno{Code: 20107, Message: "Username already exists"}
	ErrRoleExist         = &Errno{Code: 20108, Message: "Role already exists"}
	ErrNotRoleExist      = &Errno{Code: 20109, Message: "Role does not exists"}
	ErrRoleNameExist     = &Errno{Code: 20110, Message: "RoleName already exists"}

	//upload errors
	ErrUploadFile               = &Errno{Code: 20201, Message: "Error uploadFile"}
	ErrUploadMime               = &Errno{Code: 20202, Message: "Error uploadMime"}
	ErrUploadFail               = &Errno{Code: 20203, Message: "Upload fail"}
	ErrOssGenerateSignatureFail = &Errno{Code: 20204, Message: "AliyunOss fail"}
)
