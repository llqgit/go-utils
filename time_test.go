package go_utils

import (
	"testing"
	"time"
)

// 测试 timeCost
func TestTimeCost(t *testing.T) {
	defer TimeCost(time.Now())
}

func weeksInYearInWeekday(weekDay int, t *testing.T) {
	// 验证每周中的一般情况
	// weekDay [1-7]
	// 2019-12-30
	// 2019-12-31
	// 2020-01-01
	// 2020-01-02
	// 2020-01-03
	// 2020-01-04
	// 2020-01-05
	now := time.Date(2019, 12, 29+weekDay, 0, 0, 0, 0, time.Local)
	weeks, year := WeeksInYear(now)
	if weeks != 52 || year != 2019 {
		t.Logf("[%s] should be week [52] year [2019]. But got [%d] [%d]", now, weeks, year)
		t.Fail()
	}
}

// 测试 weeksInYear
func TestWeeksInYear(t *testing.T) {
	var weeks, year int

	// 验证一般情况
	now := time.Date(2019, 12, 29, 0, 0, 0, 0, time.Local)
	weeks, year = WeeksInYear(now)
	if weeks != 51 || year != 2019 {
		t.Logf("[%s] should be week [51] year [2019]. But got [%d] [%d]", now, weeks, year)
		t.Fail()
	}

	// 验证一周情况
	for i := 1; i <= 7; i++ {
		weeksInYearInWeekday(i, t)
	}

	// 验证跨年情况
	now = time.Date(2019, 12, 30, 0, 0, 0, 0, time.Local)
	weeks, year = WeeksInYear(now)
	if weeks != 52 || year != 2019 {
		t.Logf("[%s] should be week [52] year [2019]. But got [%d] [%d]", now, weeks, year)
		t.Fail()
	}
	now = time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)
	weeks, year = WeeksInYear(now)
	if weeks != 52 || year != 2019 {
		t.Logf("[%s] should be week [52] year [2019]. But got [%d] [%d]", now, weeks, year)
		t.Fail()
	}
	now = time.Date(2020, 1, 6, 0, 0, 0, 0, time.Local)
	weeks, year = WeeksInYear(now)
	if weeks != 1 || year != 2020 {
		t.Logf("[%s] should be week [1] year [2020]. But got [%d] [%d]", now, weeks, year)
		t.Fail()
	}
}

func endOfWeekCheckDay(weekDay int, t *testing.T) {
	now := time.Date(2019, 12, 29+weekDay, 0, 0, 0, 0, time.Local)
	hours := 24 * (8 - weekDay)
	shouldHave := time.Hour * time.Duration(hours)
	if r := EndOfWeek(now); r != shouldHave {
		t.Logf("[%s] should have [%s] to the end of week. But find [%s]", now, shouldHave, r)
		t.Fail()
	}
}

// 测试计算时间到周末的毫秒数（到周日）
func TestEndOfWeek(t *testing.T) {
	//var shouldHave time.Duration
	//var now time.Time
	// 验证一般情况
	// 2019-12-30 00:00:00.000000000 周一 距离周末
	// 2020-01-05 23:59:59.999999999 周日
	// 应该还有 7*24*3600 - 0.000000001 秒

	for i := 1; i <= 7; i++ {
		endOfWeekCheckDay(i, t)
	}
	//
	//// 验证周末情况
	//// 2019-12-29 00:00:00.000000000 周日 距离周末应该有 24 hour
	//now = time.Date(2019, 12, 29, 0, 0, 0, 0, time.Local)
	//shouldHave = time.Hour * 24
	//if r := EndOfWeek(now); r != shouldHave {
	//	t.Logf("[%s] should have [%s] to the end of week. But find [%s]", now, shouldHave, r)
	//	t.Fail()
	//}
	//
	//// 验证周末情况
	//// 2019-12-28 00:00:00.000000000 周日 距离周末应该有 48 hour
	//now = time.Date(2019, 12, 28, 0, 0, 0, 0, time.Local)
	//shouldHave = time.Hour * 48
	//if r := EndOfWeek(now); r != shouldHave {
	//	t.Logf("[%s] should have [%s] to the end of week. But find [%s]", now, shouldHave, r)
	//	t.Fail()
	//}

}
