package config

type GetRequest struct {
	Key string `form:"key" binding:"required"`
}

type GetCommonResponse struct {
	ID       uint64 `json:"configId"`
	ParamKey string `json:"paramKey"`
	Remark   string `json:"remark"`
	Type     int32  `json:"type"`
}

type GetResponseWithOneParam struct {
	GetCommonResponse
	ParamValue string `json:"paramValue"`
}

type GetResponseWithMultipleParams struct {
	GetCommonResponse
	ParamValue map[string]interface{} `json:"paramValue"`
}

type UpsertRequest struct {
	ID         uint64 `json:"configId"`
	ParamKey   string `json:"paramKey"`
	ParamValue string `json:"paramValue"`
	Status     int32  `json:"status"`
	Remark     string `json:"remark"`
	Type       int32  `json:"type"`
}
