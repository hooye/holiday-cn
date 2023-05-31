package holiday

import (
	"errors"
	"fmt"
	"time"
)

type GetDate struct {
	Yearmonthday time.Time
	Year         int
}

func CheckDateFormat(p *GetDate, dateStr string) error {
	// 定义期望的日期格式
	expectedLayout := "2006-01-02"

	// 解析日期字符串
	yearMonthDay, err := time.Parse(expectedLayout, dateStr)
	if err != nil {
		return errors.New("日期格式不正确，期望格式为 " + expectedLayout)
	}
	p.Yearmonthday = yearMonthDay

	p.Year = yearMonthDay.Year()

	return nil
}

func IsItHoliday(intime string) bool {

	var p GetDate
	err := CheckDateFormat(&p, intime)
	if err != nil {
		panic(err)
	}

	var year string = fmt.Sprintf("%d", p.Year)
	body := GetData(year)

	holidys := ParseJson(body)

	holidaymap := make(map[string]bool)
	for _, holiday := range holidys.Days {
		holidaymap[holiday.Date] = holiday.IsOffDay
	}

	return holidaymap[intime]
}
