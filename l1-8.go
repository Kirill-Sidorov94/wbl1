package main

import (
	"fmt"
)

func replaceBit(num int64, pos uint64, value bool) (int64, error) {
	if pos > 63 {
        return num, fmt.Errorf("pos must have beetwen 0 - 63, your pos is %d", pos)
    }

	if value {
		return (num | (1 << pos)), nil
	} else {
		return (num &^ (1 << pos)), nil
	}
}