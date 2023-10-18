package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestSortFile(t *testing.T) {
	// Создаем временную директорию для тестовых файлов
	tempDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Создаем временный файл с данными для тестирования
	tempFile, err := os.Create(filepath.Join(tempDir, "test_input.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer tempFile.Close()

	data := []byte("grape 8\nbanana 5\ncherry 3\ndate 5\napple 10\n")
	tempFile.Write(data)

	// Запускаем тесты, передавая имя временной директории и временного файла
	tests := []struct {
		name     string
		fileName string
		column   int
		numeric  bool
		reverse  bool
		unique   bool
		expected []string
	}{
		{
			name:     "SortByColumn0NumericReverseUnique",
			fileName: tempFile.Name(),
			column:   0,
			numeric:  true,
			reverse:  true,
			unique:   true,
			expected: []string{"grape 8", "date 5", "cherry 3", "banana 5", "apple 10"},
		},
		// Другие тестовые случаи здесь
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outputFileName := filepath.Join(tempDir, "sorted_"+filepath.Base(test.fileName))
			err := SortFile(test.fileName, test.column, test.numeric, test.reverse, test.unique)
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}

			// Считываем отсортированные строки из файла
			resultData, err := ioutil.ReadFile(outputFileName)
			if err != nil {
				t.Fatal(err)
			}
			resultLines := strings.Split(string(resultData), "\n")

			// Убираем последний элемент (пустую строку)
			if len(resultLines) > 0 && resultLines[len(resultLines)-1] == "" {
				resultLines = resultLines[:len(resultLines)-1]
			}

			// Проверяем, что отсортированные строки соответствуют ожидаемым результатам
			if !reflect.DeepEqual(resultLines, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, resultLines)
			}
		})
	}
}
