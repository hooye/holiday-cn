package holiday

import (
	"fmt"
	"testing"
	"time"
)

// TestGetData 函数测试
func TestGetData(t *testing.T) {
	fmt.Println(t.Name())
	var p GetDate
	now := time.Now()

	DateFormat(&p, now)

	fmt.Printf("获取的年份是: %v\n", p.Year)

}

func TestParseJson(t *testing.T) {
	fmt.Println(t.Name())
	var p GetDate
	now := time.Now()

	DateFormat(&p, now)
	body := GetData(p.Year)

	holidays := ParseJson(body)
	if len(holidays.Days) == 0 {
		fmt.Printf("当前年份还没有发布日历\n")
	}
}

func TestCheckDateFormat(t *testing.T) {
	fmt.Println(t.Name())
	var p GetDate

	now := time.Now()

	DateFormat(&p, now)

}

func TestIsItHoliday(t *testing.T) {
	fmt.Println(t.Name())
	// now := time.Now()
	now := time.Date(2023, 5, 7, 0, 0, 0, 0, time.UTC)

	if IsItHoliday(now) {
		fmt.Println("今天是假期哦")
		return
	}
	fmt.Println("放锤子个假")
}
