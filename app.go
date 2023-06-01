package holiday

import (
	"time"
)

// 获取日期
type GetDate struct {
	Date string
	Year int
}

// DateFormat 日期格式化
func DateFormat(p *GetDate, intime time.Time) {

	// 解析日期字符串
	date := intime.Format(time.DateOnly)
	p.Date = date

	p.Year, _, _ = intime.Date()

}

// IsItHoliday 判断是否是假期
func IsItHoliday(intime time.Time) bool {

	var p GetDate
	DateFormat(&p, intime)
	body := GetData(p.Year)

	holidys := ParseJson(body)

	holidaymap := make(map[string]bool)
	for _, holiday := range holidys.Days {
		holidaymap[holiday.Date] = holiday.IsOffDay
	}

	today := intime.Weekday()
	// 短变量声明
	if isHoliday, ok := holidaymap[p.Date]; ok {
		return isHoliday
	}
	if today == time.Saturday || today == time.Sunday {
		return true
	}

	return false

}
