package user

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

type CreateResponse struct {
	Username         string `json:"username"`
	Token            string `json:"token"`
	ExpiredAt        string `json:"expired_at"`
	RefreshExpiredAt string `json:"refresh_expired_at"`
}

type ListRequest struct {
	UserName string `form:"username"`
	Page     uint64 `form:"page"`
}
