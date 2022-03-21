package utils

import (
	"fmt"
	"strconv"
	"time"
)

func ConvertIntToStr(i int64, digits int) string {
	str := strconv.FormatInt(i, 10)
	if len(str) < digits {
		//return fmt.Sprintf("%012v", str)
		format := fmt.Sprintf("%s0%vv", "%", digits)
		txt := fmt.Sprintf(format, str)
		return txt
	}

	return str[:digits]
}

func ConvertStrToInt(s string) int64 {
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return id
}

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1).Add(-time.Second)
}

func BeginOfDay(days int) time.Time {
	utc := time.Now().UTC()
	//location, err := time.LoadLocation("Asia/Singapore")
	//if err != nil {
	//	log.Print(err)
	//}
	//local := utc.In(location)
	today := StartOfDay(utc)

	return today.AddDate(0, 0, days).UTC()
}
