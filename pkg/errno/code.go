package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation        = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase          = &Errno{Code: 20002, Message: "Database error."}
	ErrToken             = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrCasbin            = &Errno{Code: 20004, Message: "rbac casbin error"}
	ErrNotPermission     = &Errno{Code: 20005, Message: "rbac无权限"}
	ErrRecordNotFound    = &Errno{Code: 20006, Message: "record not found"}
	ErrRecordExist       = &Errno{Code: 20007, Message: "该条记录已存在"}
	ErrRecordHasChildren = &Errno{Code: 20008, Message: "该条记录存在子分类"}
	ErrParams            = &Errno{Code: 20012, Message: "无效参数"}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "用户不存在"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "密码错误"}
	ErrUserExist         = &Errno{Code: 20105, Message: "用户名已存在"}
	ErrNotUserExist      = &Errno{Code: 20106, Message: "无效用户"}
	ErrDisabledUser      = &Errno{Code: 20107, Message: "用户已禁用"}
	ErrRoleExist         = &Errno{Code: 20108, Message: "角色名已存在"}

	//upload errors
	ErrUploadFile               = &Errno{Code: 20201, Message: "Error uploadFile"}
	ErrUploadMime               = &Errno{Code: 20202, Message: "Error uploadMime"}
	ErrUploadFail               = &Errno{Code: 20203, Message: "Upload fail"}
	ErrOssGenerateSignatureFail = &Errno{Code: 20204, Message: "AliYunOss Signature fail"}
	ErrAliYunBucket             = &Errno{Code: 20205, Message: "阿里云OSS Bucket读取失败"}
	ErrAliYunOssUploadFail      = &Errno{Code: 20206, Message: "阿里云OSS上传失败"}
)
