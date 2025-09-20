package main

import (
	"fmt"
	"time"
	"os"
	"github.com/beevik/ntp"
)

func main() {
	time, err := ntpTime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(time)
}

func ntpTime() (time.Time, error) {
	timeNtp, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, fmt.Errorf("main.ntpTime(): %w", err)
	}

	return timeNtp, nil
}