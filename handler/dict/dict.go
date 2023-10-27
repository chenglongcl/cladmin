package dict

type CreateDictTypeRequest struct {
	DictType string `json:"dictType" binding:"required"`
	DictName string `json:"dictName" binding:"required"`
	Remark   string `json:"remark"`
	Sort     uint64 `json:"sort"`
}

type UpdateDictTypeRequest struct {
	ID       uint64 `json:"dictTypeId" binding:"required"`
	DictType string `json:"dictType" binding:"required"`
	DictName string `json:"dictName" binding:"required"`
	Remark   string `json:"remark"`
	Sort     uint64 `json:"sort"`
}

type GetDictTypeRequest struct {
	ID       uint64 `form:"id" binding:"required_without=DictType"`
	DictType string `form:"dictType" binding:"required_without=ID"`
}

type GetDictTypeResponse struct {
	ID       uint64 `json:"dictTypeId"`
	DictType string `json:"dictType"`
	DictName string `json:"dictName"`
	Remark   string `json:"remark"`
	Sort     uint64 `json:"sort"`
}

type ListDictTypeRequest struct {
	DictType string `form:"dictType"`
	DictName string `form:"dictName"`
	Page     uint64 `form:"page"`
	Limit    uint64 `form:"limit"`
}

type DeleteDictTypeRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}

type CreateDictDataRequest struct {
	DictTypeID uint64 `json:"dictTypeId" binding:"required"`
	DictLabel  string `json:"dictLabel" binding:"required"`
	DictValue  string `json:"dictValue" binding:"required"`
	Remark     string `json:"remark"`
	Sort       uint64 `json:"sort"`
}

type UpdateDictDataRequest struct {
	ID         uint64 `json:"dictDataId" binding:"required"`
	DictTypeID uint64 `json:"dictTypeId" binding:"required"`
	DictLabel  string `json:"dictLabel" binding:"required"`
	DictValue  string `json:"dictValue" binding:"required"`
	Remark     string `json:"remark"`
	Sort       uint64 `json:"sort"`
}

type GetDictDataRequest struct {
	ID uint64 `form:"id" binding:"required"`
}

type GetDictDataResponse struct {
	ID         uint64 `json:"dictDataId"`
	DictTypeID uint64 `json:"dictTypeId"`
	DictLabel  string `json:"dictLabel"`
	DictValue  string `json:"dictValue"`
	Remark     string `json:"remark"`
	Sort       uint64 `json:"sort"`
}

type ListDictDataRequest struct {
	DictTypeID uint64 `form:"dictTypeId"`
	DictLabel  string `form:"dictLabel"`
	DictValue  string `form:"dictValue"`
	Page       uint64 `form:"page"`
	Limit      uint64 `form:"limit"`
}

type DeleteDictDataRequest struct {
	Ids []uint64 `json:"ids" binding:"required"`
}
