package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words *[]string) *map[string][]string {
	anagramMap := make(map[string][]string)

	for _, word := range *words {
		// Приводим слово к нижнему регистру
		word = strings.ToLower(word)

		// Преобразуем строку в срез рун и сортируем руны
		runes := []rune(word)
		sort.Sort(RuneSlice(runes))

		// Преобразуем отсортированные руны обратно в строку
		sortedWord := string(runes)

		// Добавляем отсортированное слово в карту анаграмм
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	// Удаляем множества из одного элемента
	for sortedWord, anagrams := range anagramMap {
		if len(anagrams) == 1 {
			delete(anagramMap, sortedWord)
		}
	}

	return &anagramMap
}

// Структура и методы для сортировки рун
type RuneSlice []rune

func (rs RuneSlice) Len() int           { return len(rs) }
func (rs RuneSlice) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }
func (rs RuneSlice) Less(i, j int) bool { return rs[i] < rs[j] }

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "плитка", "каплит", "литкап", "талкип", "ежик"}
	anagramMap := findAnagrams(&words)

	// Выводим результат
	for _, anagrams := range *anagramMap {
		fmt.Printf("Множество анаграмм для %s: %v\n", anagrams[0], anagrams)
	}
}
