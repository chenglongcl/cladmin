package article

type CreateRequest struct {
	Uid     uint64   `json:"uid"`
	CateId  uint64   `json:"cate_id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Images  []string `json:"images"`
}

type CreateResponse struct {
	Id        uint64   `json:"id"`
	Uid       uint64   `json:"uid"`
	CateId    uint64   `json:"cate_id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Images    []string `json:"images"`
	Author    Author   `json:"author"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type ListRequest struct {
	CateId uint64 `json:"cate_id"`
	Page   uint64 `json:"page"`
}

type Author struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}
