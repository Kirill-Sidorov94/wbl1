package main

var justString string

func l115() {
	// Утечка так как используем срез из возможно большой строки
	// Возможность срезать символ, паника
	// v := createHugeString(1 << 10)
	// justString = v[:100]
	
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}