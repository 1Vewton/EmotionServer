package api

import (
	"github.com/gin-gonic/gin"
)

// Response defines the basic response of the response
type Response struct {
	Success bool    `json:"success"`
	Error   *string `json:"error"`
	Data    any     `json:"data"`
}

// NewResponse creates response
func NewResponse(
	c *gin.Context,
	code int,
	success bool,
	data any,
	err error,
) {
	var errInfo *string
	if err == nil {
		errInfo = nil
	} else {
		errString := err.Error()
		errInfo = &errString
	}
	newResponse := Response{
		Success: success,
		Error:   errInfo,
		Data:    data,
	}
	c.JSON(code, newResponse)
}
