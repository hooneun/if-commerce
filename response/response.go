package response

type response struct {
	Result  bool        `json:"result"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(code int, message string, data interface{}) response {
	return response{
		Result:  true,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func Err(code int, message string) response {
	return response{
		Result:  false,
		Code:    code,
		Message: message,
	}
}
