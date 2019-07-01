package menu

type CreateRequest struct {
	ParentID uint64 `json:"parentId" binding:"exists"`
	Name     string `json:"name" binding:"required"`
	Url      string `json:"url" binding:"exists"`
	Perms    string `json:"perms"`
	Type     int64  `json:"type" binding:"exists"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
}

type GetRequest struct {
	ID uint64 `form:"id"`
}

type GetResponse struct {
	ID         uint64 `json:"menuId"`
	ParentID   uint64 `json:"parentId"`
	ParentName string `json:"parentName"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Perms      string `json:"perms"`
	Type       int64  `json:"type"`
	Icon       string `json:"icon"`
	OrderNum   int64  `json:"orderNum"`
	Open       int64  `json:"open"`
	CreateTime string `json:"createTime"`
}

type UpdateRequest struct {
	ID       uint64 `json:"menuId" binding:"required"`
	ParentID uint64 `json:"parentId" binding:"exists"`
	Name     string `json:"name" binding:"required"`
	Url      string `json:"url" binding:"required"`
	Perms    string `json:"perms"`
	Type     int64  `json:"type" binding:"required"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
}

type DeleteRequest struct {
	ID uint64 `form:"id" binding:"required"`
}
