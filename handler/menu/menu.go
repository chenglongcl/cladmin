package menu

type CreateRequest struct {
	ParentId int64  `json:"parent_id" binding:"required" validate:"required"`
	Name     string `json:"name" binding:"required" validate:"required"`
	Url      string `json:"url"`
	Perms    string `json:"perms"`
	Type     int64  `json:"type"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"order_num"`
}
