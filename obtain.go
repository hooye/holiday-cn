package holiday

import (
	"fmt"
	"io"
	"net/http"
)

func GetData(year int) []byte {

	url := fmt.Sprintf("https://raw.githubusercontent.com/NateScarlet/holiday-cn/master/%v.json", year)

	// 创建 HTTP GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("无法创建 HTTP 请求：", err)
	}

	// 发送 HTTP 请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("无法发送 HTTP 请求：", err)

	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("无法读取 HTTP 响应：", err)
	}

	return body
}
