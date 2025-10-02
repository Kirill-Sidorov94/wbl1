package sorter

import (
	"strconv"
	"time"
	"strings"

	"github.com/Kirill-Sidorov94/wbl1/utils/sort/config"
)

var (
	months = map[string]time.Month{
		"jan": time.January, 
		"feb": time.February, 
		"mar": time.March,
		"apr": time.April,   
		"may": time.May,      
		"jun": time.June,
		"jul": time.July,    
		"aug": time.August,   
		"sep": time.September,
		"oct": time.October, 
		"nov": time.November, 
		"dec": time.December,
	}
	multipliers = map[string]float64{
		"k": 1e3, 
		"m": 1e6, 
		"g": 1e9, 
		"t": 1e12,
		"kb": 1e3, 
		"mb": 1e6, 
		"gb": 1e9, 
		"tb": 1e12,
	}
)

// less определеяем правило сортировки
func (s *Sorter) less(a, b string) bool {
	if s.keyColumn > 0 {
		a = getColumn(a, s.keyColumn)
		b = getColumn(b, s.keyColumn)
	}

	if s.ignoreBlanks {
		a = strings.TrimRight(a, " \t")
		b = strings.TrimRight(b, " \t")
	}

	var result bool
	switch s.sortType {
	case config.NumericSortType:
		result = lessNumeric(a, b)
	case config.MonthSortType:
		result = lessMonth(a, b)
	case config.HumanSortType:
		result = lessHuman(a, b)
	default:
		result = a < b
	}

	if s.reverse {
		return !result
	}
	return result
}

// getColumn возвращает указанную колонку
func getColumn(line string, col int) string {
	fields := strings.Fields(line)
	idx := col-1
	if idx >= 0 && idx < len(fields) {
		return fields[idx]
	}
	return ""
}

// lessNumeric сравнение чисел
func lessNumeric(a, b string) bool {
	numA, errA := strconv.ParseFloat(a, 64)
	numB, errB := strconv.ParseFloat(b, 64)

	if errA == nil && errB == nil {
		return numA < numB
	}

	return a < b
}

// lessMonth сравнение по месяцам
func lessMonth(a, b string) bool {
	monthA := parseMonth(a)
	monthB := parseMonth(b)

	if monthA != 0 && monthB != 0 {
		return monthA < monthB
	}

	return a < b
}

// parseMonth парсит месяц
func parseMonth(s string) time.Month {
	if month, ok := months[strings.ToLower(s)]; ok {
		return month
	}

	return 0
}

// lessMonth сравнение по человекочитаемым числам
func lessHuman(a, b string) bool {
	return parseHuman(a) < parseHuman(b)
}

// parseHuman парсит человекочитаемые числа (2K, 1M, 3G)
func parseHuman(s string) float64 {
	s = strings.ToLower(strings.TrimSpace(s))

	for suffix, mult := range multipliers {
		if strings.HasSuffix(s, suffix) {
			numStr := strings.TrimSuffix(s, suffix)
			num, err := strconv.ParseFloat(numStr, 64)
			if err == nil {
				return num * mult
			}
		}
	}

	// Пробуем как обычное число
	num, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return num
	}

	return 0
}
