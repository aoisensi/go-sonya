package sonya

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Raw     string `json:"-"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}
