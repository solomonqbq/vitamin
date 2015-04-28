package datetime

import "time"

var FORMAT_SECOND = "2006-01-02 15:04:05"

func NowStr() string {
	t := time.Now()
	return t.Format(FORMAT_SECOND)
}

func Parse(date_str string) (time.Time, error) {
	return time.Parse(FORMAT_SECOND, date_str)
}
