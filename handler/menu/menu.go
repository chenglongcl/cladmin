package menu

type CreateRequest struct {
	ParentId uint64 `json:"parent_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Url      string `json:"url" binding:"required"`
	Perms    string `json:"perms"`
	Type     int64  `json:"type" binding:"required"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"order_num"`
}

type GetRequest struct {
	Id uint64 `form:"id"`
}

type GetResponse struct {
	Id         uint64 `json:"id"`
	ParentId   uint64 `json:"parent_id"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Perms      string `json:"perms"`
	Type       int64  `json:"type"`
	Icon       string `json:"icon"`
	OrderNum   int64  `json:"order_num"`
	CreateTime string `json:"create_time"`
}

type UpdateRequest struct {
	Id       uint64 `json:"id" binding:"required"`
	ParentId uint64 `json:"parent_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Url      string `json:"url" binding:"required"`
	Perms    string `json:"perms"`
	Type     int64  `json:"type" binding:"required"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"order_num"`
}

type DeleteRequest struct {
	Id uint64 `form:"id" binding:"required"`
}
