package types

// ErrorResponse - Return error response
type ErrorResponse struct {
	ErrorCode      string `json:"err_code,omitempty"`
	ErrorHeading   string `json:"err_heading,omitempty"`
	ErrorMsg       string `json:"err_msg,omitempty"`
	SystemErrorMsg string `json:"sys_err_msg,omitempty"`
	Status         int    `json:"-"`
}

func (e ErrorResponse) Error() string {
	return e.ErrorMsg
}

type HTTPResponse struct {
	Data       interface{}    `json:"data,omitempty"`
	Error      *ErrorResponse `json:"error"`
	Message    string         `json:"message"`
	StatusCode int            `json:"status_code"`
}
