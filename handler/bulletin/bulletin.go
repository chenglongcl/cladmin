package bulletin

type CreateRequest struct {
	Title   string `json:"title" binding:"required"`
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type UpdateRequest struct {
	Id      uint64 `json:"id" binding:"exists"`
	Title   string `json:"title" binding:"required"`
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type GetRequest struct {
	Id uint64 `form:"id" binding:"exists"`
}

type GetResponse struct {
	Id         uint64 `json:"id"`
	Title      string `json:"title"`
	Tag        string `json:"tag"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
}

type ListRequest struct {
	Title string `form:"title"`
	Tag   string `form:"tag"`
	Page  uint64 `form:"page"`
	Limit uint64 `form:"limit"`
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}
