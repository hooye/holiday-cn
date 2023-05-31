package holiday

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Holiday struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	IsOffDay bool   `json:"isOffDay"`
}

type Holidays struct {
	// Year int       `json:"year"`
	Days []Holiday `json:"days"`
}

func ParseJson(data []byte) Holidays {
	// 解析 JSON 数据
	var holidays Holidays
	err := jsoniter.Unmarshal(data, &holidays)
	if err != nil {
		fmt.Println("无法解析 JSON 数据：", err)

	}

	return holidays
}
