package utils

const (
	StatusError   = "error"
	StatusSuccess = "success"
)

func SuccessResonse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": StatusSuccess,
		"data":   data,
	}
}

func ErrorResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  StatusError,
		"message": message,
	}
}
