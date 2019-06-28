package util

import "time"

type Interval struct {
	Ticker    *time.Ticker
	Stop      chan bool
	Start     chan bool
	IsRunning bool
}

// Ticker 실행
func CreateInterval(do func(), period time.Duration) *Interval {
	interval := Interval{time.NewTicker(period * time.Second), make(chan bool, 1), make(chan bool, 1), true}
	StopInterval(&interval)

	go func() {
		for {
			select {
			case <-interval.Ticker.C:
				do()
			case <-interval.Stop:
				interval.Ticker.Stop()
			case <-interval.Start:
				interval.Ticker = time.NewTicker(period * time.Second)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	return &interval
}

// Ticker 중지
func StopInterval(interval *Interval) {
	interval.IsRunning = false
	interval.Stop <- true
}

// Ticker 재개
func StartInterval(interval *Interval) {
	interval.IsRunning = true
	interval.Start <- true
}
