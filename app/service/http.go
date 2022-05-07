package service

import (
	"github.com/gin-gonic/gin"
	"fmt"
	// "io/ioutil"
	"net/http"
)

func Http(c *gin.Context) {
	response, err := http.Get("http://www.baidu.com")

	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)

	// fmt.Println(string(body))

	fmt.Println(response.StatusCode)

	if response.StatusCode == 200 {
		fmt.Println("ok")
	}
	c.JSON(200, gin.H{
		"code": 200,
		"message": "http",
	})
}