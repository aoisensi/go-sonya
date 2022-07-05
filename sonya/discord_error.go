package sonya

import "fmt"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Raw     string `json:"-"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("sonya: discord returned error response: #%v %v", e.Code, e.Message)
}
