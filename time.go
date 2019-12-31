package go_utils

import (
	"fmt"
	"time"
)

// 获取周数（1-7，周一到周日）(默认是 0-6)
func GetWeekDay(t time.Time) int {
	weekday := int(t.Weekday())
	// 周日使用 7（默认是 0）
	if weekday == 0 {
		weekday = 7
	}
	return weekday
}

// 计算函数消耗时间
func TimeCost(start time.Time, msg ...string) {
	terminal := time.Since(start)
	if len(msg) > 0 {
		fmt.Printf("[%v]time cost:%v\n", msg, terminal)
		return
	}
	fmt.Println("time cost:", terminal)
}

// WeeksInYear 计算时间是当年的第几周（周一为第一天）（根据本周周一判断本周是哪年的）
func WeeksInYear(today time.Time) (weeks int, year int) {
	/************ 避免跨年特殊情况 ************/
	// 使用本周第一天（周一）作为本周所处哪一周的具体判断（避免跨年时周数被撕裂的情况）
	weekday := GetWeekDay(today)
	t := today.AddDate(0, 0, -weekday)

	/************ 正常计算周数 ************/
	yearDay := t.YearDay()
	// 获取今年的年数 eg:2019
	year = t.Year()
	// 判断第一天是周几
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())
	// 今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		// 如果整除时，为周日。此时应该判断周日是每周第一天，还是每周最后一天
		base := 2
		if (yearDay-firstWeekDays)%7 == 0 {
			base = 1
		}
		week = (yearDay-firstWeekDays)/7 + base
	}
	return week, year
}

// 计算时间到周末的毫秒数（到周日）
func EndOfWeek(t time.Time) time.Duration {
	weekday := int(t.Weekday())
	// 周日使用 7（默认是 0）
	if weekday == 0 {
		weekday = 7
	}
	weekend := time.Date(t.Year(), t.Month(), t.Day()+8-weekday, 0, 0, 0, 0, t.Location()).UnixNano()
	return time.Duration(weekend - t.UnixNano())
}

// LastWeeksInYear 计算上一周是哪年的第几周
func LastWeeksInYear() (weeks int, year int) {
	t := time.Now()
	lastWeek := t.AddDate(0, 0, -7)
	return WeeksInYear(lastWeek)
}

// 明天 0 点时间戳
func NextDayZero(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, t.Location())
}

// 一天的起始时间
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// 一天的最后时间（23:59:59.999999999）
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, -1, t.Location())
}

// 获取毫秒时间戳
func Milliseconds() int64 {
	return time.Now().UnixNano() / 1e3
}

// 获取今年现在所处的天数
func DaysInYear() int {
	return time.Now().YearDay()
}
