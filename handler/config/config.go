package config

type GetRequest struct {
	Key string `form:"key" binding:"required"`
}

type GetCommonResponse struct {
	Id       uint64 `json:"configId"`
	ParamKey string `json:"paramKey"`
	Remark   string `json:"remark"`
	Type     int64  `json:"type"`
}

type GetResponseWithOneParam struct {
	GetCommonResponse
	ParamValue string `json:"paramValue"`
}

type GetResponseWithMultipleParams struct {
	GetCommonResponse
	ParamValue map[string]interface{} `json:"paramValue"`
}

type UpdateRequest struct {
	Id         uint64 `json:"configId" binding:"required"`
	ParamKey   string `json:"paramKey"`
	ParamValue string `json:"paramValue"`
	Status     int64  `json:"status"`
	Remark     string `json:"remark"`
	Type       int64  `json:"type"`
}
