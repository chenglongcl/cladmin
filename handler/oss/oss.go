package oss

type SaveConfigRequest struct {
	AliYunAccessKeyId     string `json:"aliyunAccessKeyId"`
	AliYunAccessKeySecret string `json:"aliyunAccessKeySecret"`
	AliYunBucketName      string `json:"aliyunBucketName"`
	AliYunEndPoint        string `json:"aliyunEndPoint"`
	OssType               string `json:"ossType"`
}

type UploadOssRequest struct {
	OssName string `form:"ossName" binding:"required"`
}
