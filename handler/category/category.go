package category

type CreateRequest struct {
	ParentId uint64 `json:"parentId" binding:"exists"`
	Name     string `json:"name" binding:"required"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
}

type UpdateRequest struct {
	Id       uint64 `json:"id" binding:"exists"`
	ParentId uint64 `json:"parentId" binding:"exists"`
	Name     string `json:"name" binding:"required"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
}

type GetRequest struct {
	Id uint64 `form:"id" binding:"exists"`
}

type GetResponse struct {
	Id         uint64 `json:"categoryId"`
	ParentId   uint64 `json:"parentId"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	OrderNum   int64  `json:"orderNum"`
	CreateTime string `json:"createTime"`
}

type DeleteRequest struct {
	Id uint64 `form:"id" binding:"exists"`
}
