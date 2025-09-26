package main

import (
	"math/rand"
    "strings"
    "time"
)

var justString string

func l115() {
	// Утечка так как используем срез из возможно большой строки
	// Возможность срезать символ, паника
	// v := createHugeString(1 << 10)
	// justString = v[:100]
	
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func createHugeString(size int) string {
    if size <= 0 {
        return ""
    }

    rand.Seed(time.Now().UnixNano())
    
    var builder strings.Builder
    builder.Grow(size)
    letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digits := "0123456789"
    
    for builder.Len() < size {
        char := letters[rand.Intn(len(letters))]
        builder.WriteByte(char)
        
        if rand.Intn(2) == 0 && builder.Len() < size-1 {
            multiplier := digits[rand.Intn(9)+1] // цифры 1-9 (исключаем 0)
            builder.WriteByte(multiplier)
        }
    }
    
    result := builder.String()
    if len(result) > size {
        return result[:size]
    }
    return result
}