package main

import (
	"time"
)

func customSleep(timeout time.Duration) {
	timer := time.NewTimer(timeout)
    <-timer.C
    timer.Stop()
}