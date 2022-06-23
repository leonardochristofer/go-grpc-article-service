package timestamp

import "time"

func GetNow() string {
	return time.Now().Format(time.RFC3339)
}
