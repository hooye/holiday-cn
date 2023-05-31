package holiday

import (
	"fmt"
	"testing"
)

func TestGetData(t *testing.T) {
	fmt.Println(t.Name())
	var p GetDate
	err := CheckDateFormat(&p, "2023-06-01")
	if err != nil {
		return
	}
	fmt.Printf("获取的年份是: %v\n", p.Year)

}

func TestParseJson(t *testing.T) {
	fmt.Println(t.Name())
	body := GetData("2024")
	holidays := ParseJson(body)
	if len(holidays.Days) == 0 {
		fmt.Printf("当前年份还没有发布日历\n")
	}
}

func TestCheckDateFormat(t *testing.T) {
	fmt.Println(t.Name())
	var p GetDate
	err := CheckDateFormat(&p, "2022-01-01")
	if err != nil {
		fmt.Print(err)
	}

}

func TestIsItHoliday(t *testing.T) {
	fmt.Println(t.Name())
	if IsItHoliday("2023-06-22") {
		fmt.Println("今天是假期哦")
		return
	}
	fmt.Println("放锤子个假")
}
