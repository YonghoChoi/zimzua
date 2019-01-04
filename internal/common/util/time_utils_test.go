package util

import (
	"testing"
	"time"
	nsomredis "gitlab.nexon.com/irene/nxkit/nsom/cmd/nsom-agent/pkg/redis"
	"gitlab.nexon.com/irene/nxkit/nsom/internal/common/etc"
	"os"
)

func TestNow(t *testing.T) {
	t.Log(GetCurrentTimeInSeoulToString())

	now := time.Now().UTC()
	nowStr := now.Format(time.RFC3339)
	t.Log(nowStr)
	parseTime, _ := time.Parse(time.RFC3339, nowStr)
	t.Log(parseTime)
}

func TestRedisTime(t *testing.T) {
	if os.Getenv("REDIS_ADDR") == "" {
		t.Skip("Skipped... Redis setting does not exist.")
	}

	now, err := nsomredis.GetRedisTimestampToSec()
	if err != nil {
		t.Error("GetRedisTimestampToSec fail: ", err)
	}

	time.Sleep(5 * time.Second)
	end, err := nsomredis.GetRedisTimestampToSec()
	if err != nil {
		t.Error("GetRedisTimestampToSec fail: ", err)
	}

	t.Log(end - now)

	if err := nsomredis.Command("SETEX", "test", etc.STATUS_EXPIRE_SEC, ""); err != nil {
		t.Error("status error: ", err)
	}
}

func TestTime(t *testing.T) {
	t.Log(int((etc.TIMEOUT_MINUTES_CUSTOM_PROCESS_EXECUTING_TIME * time.Minute).Seconds()))
}
