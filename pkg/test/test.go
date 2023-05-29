package test

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hooye/holiday-cn/pkg/obtain"
)

var HelperPath = "./helper/test.json"

func LocalTest() {

	fmt.Println("localttest")
	_, err := os.Stat(HelperPath)
	if os.IsNotExist(err) {
		SaveFile()
	}

}

// SaveFile 发起请求获取文件
func SaveFile() {

	data := obtain.GetData()

	file, err := os.OpenFile(HelperPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		fmt.Println(err)
	}
	writer.Flush()

}
