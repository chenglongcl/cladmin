package role

type CreateRequest struct {
	RoleName     string  `json:"role_name" binding:"required" validate:"required"`
	Remark       string  `json:"remark"`
	CreateUserId uint64  `json:"create_user_id" validate:"min=1"`
	MenuIdList   []int64 `json:"menu_id_list" binding:"required" validate:"required"`
}
