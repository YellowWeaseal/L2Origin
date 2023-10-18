package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

type FilterOptions struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
}

func main() {
	options := parseFlags()
	pattern := buildPattern(options)

	input := openInputFile(options)
	defer input.Close()

	filterText(input, pattern, options)
}

func parseFlags() *FilterOptions {
	var options FilterOptions
	flag.IntVar(&options.after, "A", 0, "Печатать +N строк после совпадения")
	flag.IntVar(&options.before, "B", 0, "Печатать +N строк до совпадения")
	flag.IntVar(&options.context, "C", 0, "Печатать ±N строк вокруг совпадения")
	flag.BoolVar(&options.count, "c", false, "Количество строк")
	flag.BoolVar(&options.ignoreCase, "i", false, "Игнорировать регистр")
	flag.BoolVar(&options.invert, "v", false, "Исключать совпадения")
	flag.BoolVar(&options.fixed, "F", false, "Точное совпадение со строкой, не паттерн")
	flag.BoolVar(&options.lineNum, "n", false, "Напечатать номер строки")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Пожалуйста, укажите паттерн для поиска.")
		os.Exit(1)
	}

	options.pattern = flag.Arg(0)

	return &options
}

func buildPattern(options *FilterOptions) string {
	pattern := options.pattern
	if options.fixed {
		pattern = regexp.QuoteMeta(pattern)
	}
	if options.ignoreCase {
		pattern = "(?i)" + pattern
	}
	return pattern
}

func openInputFile(options *FilterOptions) *os.File {
	if flag.NArg() > 1 {
		filePath := flag.Arg(1)
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Ошибка при открытии файла:", err)
			os.Exit(1)
		}
		return file
	}
	return os.Stdin
}

func filterText(input *os.File, pattern string, options *FilterOptions) {
	scanner := bufio.NewScanner(input)
	lineNumber := 0
	matchCount := 0
	beforeBuffer := []string{}
	afterBuffer := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		matched, _ := regexp.MatchString(pattern, line)
		if (options.invert && !matched) || (!options.invert && matched) {
			matchCount++

			if options.count {
				continue
			}

			if options.lineNum {
				fmt.Printf("%d: ", lineNumber)
			}

			fmt.Println(line)

			if options.context > 0 {
				for _, bline := range beforeBuffer {
					fmt.Println(bline)
				}
				beforeBuffer = []string{}
			}

			if options.after > 0 || options.context > 0 {
				for i := 0; i < options.after; i++ {
					if scanner.Scan() {
						afterLine := scanner.Text()
						if options.context > 0 {
							afterBuffer = append(afterBuffer, afterLine)
						} else {
							fmt.Println(afterLine)
						}
					} else {
						break
					}
				}
				if options.context > 0 {
					for _, aline := range afterBuffer {
						fmt.Println(aline)
					}
					afterBuffer = []string{}
				}
			}
		} else if options.before > 0 || options.context > 0 {
			if len(beforeBuffer) >= options.before {
				beforeBuffer = beforeBuffer[1:]
			}
			beforeBuffer = append(beforeBuffer, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении:", err)
	}

	if options.count {
		fmt.Printf("Количество совпадений: %d\n", matchCount)
	}
}
