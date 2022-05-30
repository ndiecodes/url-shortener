package main

import (
	"net/http"

	"url-shortener/utils"

	"github.com/gin-gonic/gin"
)

var urls = map[string]string{}

func main() {
	r := gin.Default()
	r.GET("/add", func(c *gin.Context) {
		url := c.Query("url")

		if len(url) == 0 || !utils.IsUrl(url) {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid URL"))
			return
		}
		urlIndex := utils.GetURLIndex()
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
