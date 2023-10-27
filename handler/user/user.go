package user

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=1,max=32"`
	Password string `json:"password" binding:"required,min=6,max=128"`
}

type CreateRequest struct {
	Username     string   `json:"username" binding:"required,min=1,max=32"`
	Password     string   `json:"password" binding:"required,min=6,max=128"`
	DeptID       uint64   `json:"deptId" binding:"required"`
	Mobile       string   `json:"mobile" binding:"numeric,max=11"`
	Email        string   `json:"email" binding:"email"`
	Gender       int32    `json:"gender"`
	Status       int32    `json:"status"`
	CreateUserID uint64   `json:"createUserId"`
	RoleIdList   []uint64 `json:"roleIdList"`
}

type CreateResponse struct {
	Username         string `json:"username"`
	Token            string `json:"token"`
	ExpiredAt        string `json:"expired_at"`
	RefreshExpiredAt string `json:"refresh_expired_at"`
}

type ListRequest struct {
	UserName string `form:"username"`
	Page     uint64 `form:"page"`
	Limit    uint64 `form:"limit"`
}

type GetRequest struct {
	ID uint64 `form:"id"`
}

type GetResponse struct {
	UserID       uint64   `json:"userId"`
	Username     string   `json:"username"`
	CreateTime   string   `json:"createTime"`
	CreateUserID uint64   `json:"createUserId"`
	DeptID       uint64   `json:"deptId"`
	DeptName     string   `json:"deptName"`
	Email        string   `json:"email"`
	Mobile       string   `json:"mobile"`
	Gender       int32    `json:"gender"`
	SuperAdmin   bool     `json:"superAdmin"`
	Status       int32    `json:"status"`
	RoleIDList   []uint64 `json:"roleIdList"`
}

type UpdateRequest struct {
	ID         uint64   `json:"userId" binding:"required"`
	Username   string   `json:"username" binding:"required,min=1,max=32"`
	Password   string   `json:"password"`
	DeptID     uint64   `json:"deptId" binding:"required"`
	Mobile     string   `json:"mobile" binding:"numeric,max=11"`
	Email      string   `json:"email" binding:"email"`
	Gender     int32    `json:"gender"`
	Status     int32    `json:"status"`
	RoleIDList []uint64 `json:"roleIdList" binding:"required"`
}

type UpdatePersonalRequest struct {
	Password string `json:"password"`
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}

type LogoutUserRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}
