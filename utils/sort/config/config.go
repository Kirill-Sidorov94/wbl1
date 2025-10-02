package config

import (
	"os"
	"strconv"
	"strings"
	"fmt"
	"errors"
)

const (
	NumericSortType = "numeric"
	MonthSortType = "month"
	HumanSortType = "human"
)

// Config конфигурация
type Config struct {
	KeyColumn    int    // -k N
	SortType     string // тип сортировки: -n, -M, -h
	Reverse      bool   // -r
	Unique       bool   // -u
	IgnoreBlanks bool   // -b
	CheckSorted  bool   // -c
	File         string
}

// New возвращает конфигурацию
func New() (*Config, error) {
	config := Config{}
	args := os.Args[1:] // все кроме вызова программы

	var i int
	for i < len(args) {
		arg := args[i]
		
		if len(arg) == 0 {
			i++
			continue
		}

		// Считаем файлом
		if !strings.HasPrefix(arg, "-") {
		    config.File = arg
		    i++
		    continue
		}

		// Парсим ["-k", "N"] ожидаем следующим число
		if arg == "-k" {
			if i+1 < len(args) {
				col, err := strconv.Atoi(args[i+1])
				if err != nil || col < 1 {
					return nil, fmt.Errorf("invalid number for -k: '%s'", args[i+1])
				}
				config.KeyColumn = col
				i += 2
				continue
			}

			return nil, errors.New("after -k the number N is expected")
		}

		// Обрабатываем все остальное
		for _, b := range arg[1:] {
			switch b {
			case 'n':
				config.SortType = NumericSortType
			case 'M':
				config.SortType = MonthSortType
			case 'h':
				config.SortType = HumanSortType
			case 'r':
				config.Reverse = true
			case 'u':
				config.Unique = true
			case 'b':
				config.IgnoreBlanks = true
			case 'c':
				config.CheckSorted = true
			}
		}
		i++
	}
	
	return &config, nil
}
