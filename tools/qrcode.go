package tools

import (
	"fmt"
	"image/color"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

func CreateQrcode(content string) (qrcodeUrl string, err error) {
	if content == "" {
		content = "hello world"
	}

	pngName := fmt.Sprintf("/t/%d.png", time.Now().UnixNano())
	qrcodeUrl = fmt.Sprintf("http://127.0.0.1:2020/s%s", pngName)
	// content 内容
	// Transparent 透明背景
	// White 白色字体
	err = qrcode.WriteColorFile(content, qrcode.Medium, 256, color.Transparent, color.White, "./static"+pngName)
	return
}
