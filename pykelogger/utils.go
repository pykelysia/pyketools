package pykelogger

import "time"

func getTimeStamp() string {
	return time.Now().Format("2006-02-06 13:01:02")
}
