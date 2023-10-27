package dept

type CreateRequest struct {
	ParentID uint64 `json:"parentId"`
	Name     string `json:"name" binding:"required"`
	Sort     int32  `json:"sort"`
}

type UpdateRequest struct {
	ID       uint64 `json:"deptId" binding:"required"`
	ParentID uint64 `json:"parentId"`
	Name     string `json:"name" binding:"required"`
	Sort     int32  `json:"sort"`
}

type GetRequest struct {
	ID uint64 `form:"id"`
}

type GetResponse struct {
	ID       uint64 `json:"deptId"`
	ParentID uint64 `json:"parentId"`
	Name     string `json:"name"`
	Sort     int32  `json:"sort"`
}

type DeleteRequest struct {
	ID uint64 `form:"id" binding:"required"`
}

type ListRequest struct {
}
