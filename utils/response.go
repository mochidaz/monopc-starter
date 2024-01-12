package utils

func ErrorApiResponse(code int, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": message,
	}
}

func SuccessApiResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    data,
	}
}
