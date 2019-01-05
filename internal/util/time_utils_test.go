package util

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Log(GetCurrentTimeInSeoulToString())

	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)
	t.Log(nowStr)
	parseTime, _ := time.Parse(time.RFC3339, nowStr)
	t.Log(parseTime)
}
