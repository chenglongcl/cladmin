package upload

type ImageResponse struct {
	Path     string `json:"path"`
	FileName string `json:"file_name"`
}

type UploadOssRequest struct {
	OssName string `form:"ossName" binding:"required"`
}
