package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"

	"url-shortner/utils"

	"github.com/gin-gonic/gin"
)

func getURLIndex() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

var urls = map[string]string{}

func main() {
	r := gin.Default()
	r.GET("/add", func(c *gin.Context) {
		url := c.Query("url")

		if len(url) == 0 || !IsUrl(url) {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("Invalid URL"))
			return
		}
		urlIndex := getURLIndex()
		if _, ok := urls[urlIndex]; !ok {
			urls[urlIndex] = url
		}
		c.JSON(http.StatusOK, utils.SuccessResonse(map[string]interface{}{
			"index": urlIndex,
		}))
	})

	r.GET("/:index", func(c *gin.Context) {
		urlIndex := c.Param("index")
		if _, ok := urls[urlIndex]; !ok {
			c.JSON(http.StatusNotFound, utils.ErrorResponse("URL not found"))
			return
		}
		c.JSON(http.StatusOK, utils.SuccessResonse(map[string]interface{}{
			"url": urls[urlIndex],
		}))
	})

	r.Run()
}
