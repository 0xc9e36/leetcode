package main

import (
	"fmt"
	"service/common/railchina"
	"time"
)

func main() {
	y := nightCountDown()
	fmt.Println(y)
}

func nightCountDown() int {
	t, day := railchina.ChinaTime(), railchina.ChinaTime().Day()
	nightStart, nightEnd := 23, 6

	// 22 - 06 点  如果在 23 点需要跨天, 如果是 01点则不用
	if nightStart > nightEnd && t.Hour() >= nightStart {
		day++
	}

	fmt.Println(day)

	finishTime := time.Date(t.Year(), t.Month(), day, nightEnd, 0, 0, 0, railchina.TimeZone())
	return int(finishTime.Sub(t).Seconds())
}