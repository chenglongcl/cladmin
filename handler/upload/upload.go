package upload

type ObjectInfo struct {
	Key        string
	UploadId   string
	RequestUrl string
	Location   string
}

type InitiateMultipartUploadResponse struct {
	Key        string `xml:"Key"`
	UploadId   string `xml:"UploadId"`
	RequestUrl string `xml:"RequestUrl"`
	Location   string `xml:"Location"`
}

type MultipartUploadPartRequest struct {
	PartNumber uint64 `form:"partNumber" binding:"required"`
	UploadId   string `form:"uploadId" binding:"required"`
}

type CompleteMultipartUploadRequest struct {
	CompleteMultipartUpload struct {
		Parts []Part `xml:"Part"`
	} `xml:"CompleteMultipartUpload"`
}

type ListPartsResponse struct {
	ListPartsResult ListPartsResult `xml:"ListPartsResult"`
}

type ListPartsResult struct {
	Key      string `xml:"Key"`
	UploadId string `xml:"UploadId"`
	Part     []Part `xml:"Part"`
}

type Part struct {
	PartNumber   uint64 `xml:"PartNumber"`
	LastModified string `xml:"LastModified"`
	ETag         string `xml:"ETag"`
	Size         int64  `xml:"Size"`
}

type CompleteMultipartUploadResponse struct {
	CompleteMultipartUploadResult CompleteMultipartUploadResult `xml:"CompleteMultipartUploadResult"`
}

type CompleteMultipartUploadResult struct {
	Location string `xml:"Location"`
	Key      string `xml:"Key"`
	ETag     string `xml:"ETag"`
}
