package utils

import (
	"crypto/rand"
	"fmt"
	"net/url"
)

/** JSEND */
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

/** END JSEND */

func GetURLIndex() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
