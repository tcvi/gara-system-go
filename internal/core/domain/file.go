package domain

type FileUpload struct {
	Key         string `json:"key"`
	Url         string `json:"url"`
	ContentType string `json:"content-type"`
}
