package article

type CreateRequest struct {
	UserId  uint64 `json:"userId" binding:"required"`
	CateId  uint64 `json:"cateId" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Thumb   string `json:"thumb"`
}

type UpdateRequest struct {
	Id          uint64 `json:"articleId" binding:"required"`
	UserId      uint64 `json:"userId" binding:"required"`
	CateId      uint64 `json:"cateId" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content"`
	Thumb       string `json:"thumb"`
	ReleaseTime string `json:"releaseTime"`
}

type GetRequest struct {
	Id uint64 `form:"id" binding:"exists"`
}

type GetResponse struct {
	Id          uint64 `json:"articleId"`
	CateId      uint64 `json:"cateId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Thumb       string `json:"thumb"`
	ReleaseTime string `json:"releaseTime"`
}

type ListRequest struct {
	Title  string `form:"title"`
	CateId uint64 `form:"cateId"`
	Page   uint64 `form:"page"`
	Limit  uint64 `form:"limit"`
}

type DeleteRequest struct {
	Id uint64 `form:"id" binding:"exists"`
}
