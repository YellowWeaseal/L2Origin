package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Парсим аргументы командной строки
	column := flag.Int("k", 0, "Колонка для сортировки (по умолчанию 0 - вся строка)")
	numeric := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	fileName := flag.Arg(0)

	if err := SortFile(fileName, *column, *numeric, *reverse, *unique); err != nil {
		fmt.Println("Ошибка:", err)
	}
}

// SortFile сортирует строки в файле в соответствии с параметрами
func SortFile(fileName string, column int, numeric, reverse, unique bool) error {
	// Открываем файл для чтения
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Считываем строки из файла
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Определяем функцию для сравнения строк в соответствии с параметрами сортировки
	comparator := func(i, j int) bool {
		fieldsI := strings.Fields(lines[i])
		fieldsJ := strings.Fields(lines[j])
		columnI := fieldsI[column]
		columnJ := fieldsJ[column]

		if numeric {
			numI, errI := strconv.Atoi(columnI)
			numJ, errJ := strconv.Atoi(columnJ)

			if errI == nil && errJ == nil {
				columnI = strconv.Itoa(numI)
				columnJ = strconv.Itoa(numJ)
			}
		}

		if columnI < columnJ {
			return true
		}
		if columnI > columnJ {
			return false
		}

		return lines[i] < lines[j]
	}

	// Выполняем сортировку
	sort.SliceStable(lines, func(i, j int) bool {
		if reverse {
			return !comparator(i, j)
		}
		return comparator(i, j)
	})

	// Если указан ключ -u, удаляем повторяющиеся строки
	if unique {
		var uniqueLines []string
		seen := make(map[string]bool)
		for _, line := range lines {
			if !seen[line] {
				uniqueLines = append(uniqueLines, line)
				seen[line] = true
			}
		}
		lines = uniqueLines
	}

	// Открываем файл для записи
	outputFileName := "sorted_" + filepath.Base(fileName)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Записываем отсортированные строки в файл
	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	fmt.Println("Сортировка завершена. Результат записан в", outputFileName)
	return nil
}
