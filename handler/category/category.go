package category

type CreateRequest struct {
	ParentID uint64 `json:"parentId" binding:"omitempty,number,min=0"`
	Name     string `json:"name" binding:"required"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
}

type UpdateRequest struct {
	ID       uint64 `json:"categoryId" binding:"required"`
	ParentID uint64 `json:"parentId" binding:"omitempty,number,min=0"`
	Name     string `json:"name" binding:"required"`
	Icon     string `json:"icon"`
	OrderNum int64  `json:"orderNum"`
}

type GetRequest struct {
	ID uint64 `form:"id" binding:"omitempty,number,min=0"`
}

type GetResponse struct {
	ID         uint64 `json:"categoryId"`
	ParentID   uint64 `json:"parentId"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	OrderNum   int64  `json:"orderNum"`
	CreateTime string `json:"createTime"`
}

type DeleteRequest struct {
	ID uint64 `form:"id" binding:"required"`
}
