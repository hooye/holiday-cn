package holiday

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

// Holiday Days 所有数据
type Holiday struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	IsOffDay bool   `json:"isOffDay"`
}

// Holidays data 中的所有数据
type Holidays struct {
	Days []Holiday `json:"days"`
}

// 解析获取到的 json 数据
func ParseJson(data []byte) Holidays {
	// 解析 JSON 数据
	var holidays Holidays
	err := jsoniter.Unmarshal(data, &holidays)
	if err != nil {
		fmt.Println("无法解析 JSON 数据：", err)

	}

	return holidays
}
