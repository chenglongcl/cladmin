package oss

type SaveConfigRequest struct {
	AliYunAccessKeyID     string `json:"aliyunAccessKeyId"`
	AliYunAccessKeySecret string `json:"aliyunAccessKeySecret"`
	AliYunBucketName      string `json:"aliyunBucketName"`
	AliYunEndPoint        string `json:"aliyunEndPoint"`
	OSSType               string `json:"ossType"`
}

type UploadOssRequest struct {
	OSSName string `form:"ossName" binding:"required"`
}
