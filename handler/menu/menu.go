package menu

type CreateRequest struct {
	ParentId uint64 `json:"parentId" binding:"exists"`
	Name     string `json:"name" binding:"required"`
	Url      string `json:"url" binding:"exists"`
	Perms    string `json:"perms"`
	Type     int64  `json:"type" binding:"exists"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
}

type GetRequest struct {
	Id uint64 `form:"id"`
}

type GetResponse struct {
	Id         uint64 `json:"menuId"`
	ParentId   uint64 `json:"parentId"`
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
	Id       uint64 `json:"menuId" binding:"required"`
	ParentId uint64 `json:"parentId" binding:"exists"`
	Name     string `json:"name" binding:"required"`
	Url      string `json:"url" binding:"required"`
	Perms    string `json:"perms"`
	Type     int64  `json:"type" binding:"required"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
}

type DeleteRequest struct {
	Id uint64 `form:"id" binding:"required"`
}
