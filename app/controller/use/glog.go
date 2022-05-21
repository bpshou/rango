package use

import (
	"rango/app/controller"

	"github.com/gin-gonic/gin"

	"flag"

	"github.com/golang/glog"
)

type Glog struct {
	controller.Base
}

// 使用 glog 时，需要在命令行启动时添加 `-log_dir=log -alsologtostderr` 参数
// 日志将会同时打印在 log/ 目录和标准错误输出中 -alsologtostderr
// 在 Kuberntes 中，glog 是默认日志库
func (this Glog) Glog(c *gin.Context) {
	flag.Parse()
	defer glog.Flush()

	glog.Info("hello, glog")
	glog.Warning("warning glog")
	glog.Error("error glog")

	glog.Infof("info %d", 1)
	glog.Warningf("warning %d", 2)
	glog.Errorf("error %d", 3)

	glog.V(3).Infoln("info with v 3")
	glog.V(2).Infoln("info with v 2")
	glog.V(1).Infoln("info with v 1")
	glog.V(0).Infoln("info with v 0")

	c.JSON(200, gin.H{
		"code":    200,
		"message": "Google glog put content success !",
	})
}
