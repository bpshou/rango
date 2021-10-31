package service

import (
	// "fmt"
	"log"
	"io/ioutil"
	"os/exec"
	"github.com/gin-gonic/gin"
)

func Command(c *gin.Context) {
	// command := exec.Command("ping", "www.baidu.com")
	// err := command.Run()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	cmd := exec.Command("ping", "www.baidu.com")

	stdout, err := cmd.StdoutPipe()
	if err != nil {     //获取输出对象，可以从该对象中读取输出结果
		log.Fatal(err)
	}
	defer stdout.Close()   // 保证关闭输出流

	if err := cmd.Start(); err != nil {   // 运行命令
		log.Fatal(err)
	}

	if opBytes, err := ioutil.ReadAll(stdout); err != nil {  // 读取输出结果    
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}

	c.JSON(200, gin.H{
		"code": 200,
		"message": "index",
		"res": err,
	})
}

