package article

type CreateRequest struct {
	UserID  uint64 `json:"userId" binding:"required"`
	CateID  uint64 `json:"cateId" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Thumb   string `json:"thumb"`
}

type UpdateRequest struct {
	ID          uint64 `json:"articleId" binding:"required"`
	UserID      uint64 `json:"userId" binding:"required"`
	CateID      uint64 `json:"cateId" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content"`
	Thumb       string `json:"thumb"`
	ReleaseTime string `json:"releaseTime"`
}

type GetRequest struct {
	ID uint64 `form:"id" binding:"exists"`
}

type GetResponse struct {
	ID          uint64 `json:"articleId"`
	CateID      uint64 `json:"cateId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Thumb       string `json:"thumb"`
	ReleaseTime string `json:"releaseTime"`
}

type ListRequest struct {
	Title  string `form:"title"`
	CateID uint64 `form:"cateId"`
	Page   uint64 `form:"page"`
	Limit  uint64 `form:"limit"`
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}
