package api

// Response defines the basic response of the response
type Response struct {
	Success bool    `json:"success"`
	Error   *string `json:"error"`
	Data    any     `json:"data"`
}

// NewResponse creates response
func NewResponse(
	success bool,
	data any,
	err error,
) Response {
	var errInfo *string
	if err == nil {
		errInfo = nil
	} else {
		errString := err.Error()
		errInfo = &errString
	}
	return Response{
		Success: success,
		Error:   errInfo,
		Data:    data,
	}
}
