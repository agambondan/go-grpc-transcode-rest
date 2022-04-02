package model

import "github.com/agambondan/web-go-blog-grpc-rest/app/lib"

type BaseResponse struct {
	Status     *bool   `json:"status,omitempty"`
	StatusCode *int    `json:"status_code,omitempty"`
	Message    *string `json:"message,omitempty"`
	Error      *string `json:"error,omitempty"`
}

func (br *BaseResponse) Success(message string) {
	br.Status = lib.BoolPtr(true)
	br.StatusCode = lib.Intptr(200)
	br.Message = lib.Strptr(message)
	br.Error = nil
}

func (br *BaseResponse) Failed(error, message string, statusCode int) {
	br.Status = lib.BoolPtr(true)
	br.StatusCode = lib.Intptr(statusCode)
	br.Message = lib.Strptr(message)
	br.Error = lib.Strptr(error)
}

func (br *BaseResponse) ConvertToMap() map[string]interface{} {
	result := make(map[string]interface{})
	lib.Merge(br, &result)
	return result
}
