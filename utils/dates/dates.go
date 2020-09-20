package dates

import (
	"time"
)

const (
	DateLayout   = "2006-01-02T15:04:05Z"
	JoinedLayout = "Jan 02 2006 3:04:05 PM"
	MySqlLayout  = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(DateLayout)
}

func GetDateFromString(s string) string {
	d, err := time.Parse(MySqlLayout, s)
	if err != nil {
		return "Unknown"
	}
	return d.Format(JoinedLayout)
}
