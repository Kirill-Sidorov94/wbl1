package main

func revertWordInSentence(str string) string {
	runes := []rune(str)
	reverse(runes)

	var start int
	for i := 0; i < len(runes); i++ {
		if (runes[i] == ' ' && i > 0) {
			reverse(runes[start:i])
			start = i + 1
			continue
		}

		if i == len(runes) - 1 && runes[i] != ' ' {
			reverse(runes[start:i+1])
		}
	}

	return string(runes)
}

func reverse(runes []rune) {
	start, end := 0, len(runes) - 1
	for start <= end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}