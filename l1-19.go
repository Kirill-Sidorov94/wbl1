package main

func stringReversal(str string) string {
	runes := []rune(str)
	start, end := 0, len(runes) - 1
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	return string(runes)
}