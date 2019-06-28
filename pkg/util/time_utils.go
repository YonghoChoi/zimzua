package util

import (
	"fmt"
	"strings"
	"time"
)

func NowUtcMs() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Millisecond)
}

func NowUtcSec() int64 {
	return time.Now().UTC().Unix()
}

func TimeToMs(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func NowUtcString() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func UtcStringToTime(str string) time.Time {
	parseTime, _ := time.Parse(time.RFC3339, str)
	return parseTime
}

func TimeToUtcMs(t time.Time) int64 {
	return t.UTC().UnixNano() / int64(time.Millisecond)
}

func ElapsedSec(start time.Time) int64 {
	return (NowUtcMs() - TimeToUtcMs(start)) / int64(1000)
}

func TimeSince(start time.Time) string {
	since := time.Since(start).String()
	return strings.SplitN(since, ".", 2)[0] + "s"
}

func GetCurrentTimeInSeoulToString() string {
	t := time.Now().UTC().Add(9 * time.Hour)
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
