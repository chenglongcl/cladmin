package user

type CreateRequest struct {
	Username     string  `json:"username" binding:"required" validate:"min=1,max=32"`
	Password     string  `json:"password" binding:"required" validate:"min=6,max=128"`
	Mobile       string  `json:"mobile" validate:"numeric,max=11"`
	Email        string  `json:"email"`
	Status       int64   `json:"status"`
	CreateUserId uint64  `json:"createUserId"`
	RoleIdList   []int64 `json:"roleIdList"`
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
	Id uint64 `form:"id"`
}

type GetResponse struct {
	UserId       uint64   `json:"userId"`
	Username     string   `json:"username"`
	CreateTime   string   `json:"createTime"`
	CreateUserId uint64   `json:"createUserId"`
	Email        string   `json:"email"`
	Mobile       string   `json:"mobile"`
	Status       int64    `json:"status"`
	RoleIdList   []uint64 `json:"roleIdList"`
}

type UpdateRequest struct {
	Id         uint64  `json:"userId" binding:"required"`
	Username   string  `json:"username" binding:"required"`
	Password   string  `json:"password"`
	Mobile     string  `json:"mobile" validate:"numeric,max=11"`
	Email      string  `json:"email" validate:"email"`
	Status     int64   `json:"status"`
	RoleIdList []int64 `json:"roleIdList" binding:"required"`
}

type UpdatePersonalRequest struct {
	Password string `json:"password"`
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}
