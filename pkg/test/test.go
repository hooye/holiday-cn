package test

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/hooye/holiday-cn/pkg/obtain"
	"github.com/hooye/holiday-cn/pkg/parse"
)

var HelperPath = "./helper/test.json"

func LocalTest() {

	_, err := os.Stat(HelperPath)
	if os.IsNotExist(err) {
		SaveFile()
	}

	data := obtain.GetData()
	holidays := parse.ParseJson(data)

	holidaymap := make(map[string]bool)

	//遍历节假日列表
	for _, holiday := range holidays.Days {
		// fmt.Printf("% s (% s) % v\n", holiday.Name, holiday.Date, holiday.IsOffDay)
		holidaymap[holiday.Date] = holiday.IsOffDay
	}

	now := time.Now()
	nowday := now.Format("2006-01-02")
	fmt.Printf("日期:%s", nowday)
	if holidaymap[nowday] {
		fmt.Println("今天休息哦")
		return
	}
	fmt.Println(" 上班日")

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
