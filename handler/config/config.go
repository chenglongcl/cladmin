package config

type GetRequest struct {
	ID  uint64 `form:"id" binding:"required_without=Key"`
	Key string `form:"key" binding:"required_without=ID"`
}

type GetCommonResponse struct {
	ID       uint64 `json:"configId"`
	ParamKey string `json:"paramKey"`
	Remark   string `json:"remark"`
	Type     int32  `json:"type"`
	Status   bool   `json:"status"`
}

type GetResponseWithOneParam struct {
	GetCommonResponse
	ParamValue string `json:"paramValue"`
}

type GetResponseWithMultipleParams struct {
	GetCommonResponse
	ParamValue interface{} `json:"paramValue"`
}

type CreateRequest struct {
	ParamKey   string `json:"paramKey" binding:"required"`
	ParamValue string `json:"paramValue" binding:"required"`
	Status     bool   `json:"status"`
	Remark     string `json:"remark"`
	Type       int32  `json:"type"`
}

type UpdateRequest struct {
	ID         uint64 `json:"configId" binding:"required"`
	ParamKey   string `json:"paramKey" binding:"required"`
	ParamValue string `json:"paramValue" binding:"required"`
	Status     bool   `json:"status"`
	Remark     string `json:"remark"`
	Type       int32  `json:"type"`
}

type ListRequest struct {
	ParamKey string `form:"paramKey"`
	Page     uint64 `form:"page"`
	Limit    uint64 `form:"limit"`
}

type DeleteRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}

type UpsertRequest struct {
	ID         uint64 `json:"configId"`
	ParamKey   string `json:"paramKey"`
	ParamValue string `json:"paramValue"`
	Status     bool   `json:"status"`
	Remark     string `json:"remark"`
	Type       int32  `json:"type"`
}
