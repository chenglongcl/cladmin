package role

type CreateRequest struct {
	RoleName     string   `json:"roleName" binding:"required" validate:"min=1"`
	Remark       string   `json:"remark"`
	CreateUserID uint64   `json:"createUserId" validate:"min=1"`
	MenuIDList   []uint64 `json:"menuIdList" binding:"required"`
}

type UpdateRequest struct {
	ID         uint64   `json:"roleId" binding:"required"`
	RoleName   string   `json:"roleName" binding:"required" validate:"min=1"`
	Remark     string   `json:"remark"`
	MenuIDList []uint64 `json:"menuIdList" binding:"required"`
}

type GetRequest struct {
	ID uint64 `form:"id"`
}

type GetResponse struct {
	ID           uint64   `json:"roleId"`
	RoleName     string   `json:"roleName"`
	Remark       string   `json:"remark"`
	CreateUserID uint64   `json:"createUserId"`
	MenuIDList   []uint64 `json:"menuIdList"`
	CreateTime   string   `json:"createTime"`
}

type ListRequest struct {
	RoleName string `form:"roleName"`
	Page     uint64 `form:"page"`
	Limit    uint64 `form:"limit"`
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}
