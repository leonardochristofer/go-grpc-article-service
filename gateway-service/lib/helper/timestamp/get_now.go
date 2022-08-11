package timestamp

import "time"

func GetNow() string {
	return time.Now().Format(time.RFC3339)
}

func GetNowUTC() time.Time {
	return time.Now().UTC()
}

func DateToString(t time.Time) string {
	return t.Format(time.RFC3339)
}
