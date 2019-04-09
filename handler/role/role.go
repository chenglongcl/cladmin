package role

type CreateRequest struct {
	RoleName     string  `json:"roleName" binding:"required" validate:"min=1"`
	Remark       string  `json:"remark"`
	CreateUserId uint64  `json:"createUserId" validate:"min=1"`
	MenuIdList   []int64 `json:"menuIdList" binding:"required"`
}

type UpdateRequest struct {
	Id         uint64  `json:"roleId" binding:"required"`
	RoleName   string  `json:"roleName" binding:"required" validate:"min=1"`
	Remark     string  `json:"remark"`
	MenuIdList []int64 `json:"menuIdList" binding:"required"`
}

type GetRequest struct {
	Id uint64 `form:"id"`
}

type GetResponse struct {
	Id           uint64  `json:"roleId"`
	RoleName     string  `json:"roleName"`
	Remark       string  `json:"remark"`
	CreateUserId uint64  `json:"createUserId"`
	MenuIdList   []int64 `json:"menuIdList"`
	CreateTime   string  `json:"createTime"`
}

type ListRequest struct {
	RoleName string `form:"roleName"`
	Page     uint64 `form:"page"`
	Limit    uint64 `form:"limit"`
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}
