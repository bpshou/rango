package service

import (
	"io/ioutil"
	"os/exec"
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Douyin struct {
	Downloaddir string `form:"downloaddir" json:"downloaddir" xml:"downloaddir"`
	Resource string `form:"resource" json:"resource" xml:"resource" binding:"required"`
}

func Command(c *gin.Context) {
	// 获取参数
	var json Douyin
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 200,
			"message": err.Error(),
		})
		return
	}

	Download := ""
	Resource := ""

	if json.Resource == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 200,
			"message": "resource error",
		})
		return
	} else {
		Resource = json.Resource
	}

	if json.Downloaddir == "" {
		Download = "C:\\Users\\origin\\Downloads"
	} else {
		Download = json.Downloaddir
	}

	log.WithFields(log.Fields{
		"Download": Download,
		"Resource": Resource,
	}).Info("Douyin data")

	execCommand(Download, Resource)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "Douyin download",
	})
}

func execCommand(Download string, Resource string) {
	cmd := exec.Command("annie", "-o", Download, Resource)

	stdout, err := cmd.StdoutPipe()
	if err != nil { // 获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	defer stdout.Close() // 保证关闭输出流

	if err := cmd.Start(); err != nil { // 运行命令
		log.Fatal(err)
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		log.Fatal(err)
	} else {
		log.Debug(string(opBytes))
	}
}
