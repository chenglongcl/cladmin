package role

type CreateRequest struct {
	RoleName     string  `json:"role_name" binding:"required" validate:"min=1"`
	Remark       string  `json:"remark"`
	CreateUserId uint64  `json:"create_user_id" validate:"min=1"`
	MenuIdList   []int64 `json:"menu_id_list" binding:"required"`
}

type UpdateRequest struct {
	Id         uint64  `json:"id" binding:"required"`
	RoleName   string  `json:"role_name" binding:"required" validate:"min=1"`
	Remark     string  `json:"remark"`
	MenuIdList []int64 `json:"menu_id_list" binding:"required"`
}

type GetRequest struct {
	Id uint64 `form:"id"`
}

type GetResponse struct {
	Id           uint64   `json:"id"`
	RoleName     string   `json:"role_name"`
	Remark       string   `json:"remark"`
	CreateUserId uint64   `json:"create_user_id"`
	MenuIdList   []uint64 `json:"menu_id_list"`
	CreateTime   string   `json:"create_time"`
}

type ListRequest struct {
	RoleName string `form:"role_name"`
	Page     uint64 `form:"page"`
	Limit    uint64 `form:"limit"`
}

type DeleteRequest struct {
	Id uint64 `form:"id" binding:"required"`
}
