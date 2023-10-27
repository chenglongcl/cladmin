package menu

type CreateRequest struct {
	ParentID uint64 `json:"parentId"`
	Name     string `json:"name" binding:"required"`
	URL      string `json:"url" binding:"omitempty"`
	Perms    string `json:"perms"`
	Type     int64  `json:"type" binding:"oneof=0 1 2"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
	IsTab    bool   `json:"isTab"`
	Status   bool   `json:"status"`
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
	IsTab      bool   `json:"isTab"`
	CreateTime string `json:"createTime"`
	Status     bool   `json:"status"`
}

type UpdateRequest struct {
	ID       uint64 `json:"menuId" binding:"required"`
	ParentID uint64 `json:"parentId"`
	Name     string `json:"name" binding:"required"`
	URL      string `json:"url" binding:"required_with=Type"`
	Perms    string `json:"perms"`
	Type     int64  `json:"type" binding:"oneof=0 1 2"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
	IsTab    bool   `json:"isTab"`
	Status   bool   `json:"status"`
}

type DeleteRequest struct {
	ID uint64 `form:"id" binding:"required"`
}

type ListRequest struct {
	MenuTypes string `form:"menuTypes"`
}
