package helper

import (
	"time"
)

func GetNext30MinTime() time.Time {
	t := time.Now().Add(time.Minute * 10)
	t = t.Round(time.Minute * 30)
	if t.Before(time.Now()) {
		return t.Add(time.Minute * 30)
	}
	return t
}
