package lib

import (
	"fmt"
	"time"
	"github.com/beevik/ntp"
)

// NtpTime - ищи вызов в корневом main.go l2.8
func NtpTime() (time.Time, error) {
	timeNtp, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, fmt.Errorf("main.ntpTime(): %w", err)
	}

	return timeNtp, nil
}
