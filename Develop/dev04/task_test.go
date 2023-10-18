package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "плитка", "каплит", "литкап", "талкип", "ежик"}
	expectedResult := map[string][]string{
		"акптя":  {"пятак", "пятка", "тяпка"},
		"иклост": {"листок", "слиток", "столик"},
		"аилкпт": {"плитка", "каплит", "литкап"},
		"аиклпт": {"талкип"},
	}

	result := findAnagrams(&words)

	if !reflect.DeepEqual(result, &expectedResult) {
		t.Errorf("Ожидалось %v, получено %v", &expectedResult, result)
	}
}

func TestRuneSort(t *testing.T) {
	r := RuneSlice([]rune("пятка"))
	sort.Sort(r)
	sortedWord := string(r)
	if sortedWord != "акптя" {
		t.Errorf("Ожидалась 'акптя', получено '%s'", sortedWord)
	}
}

func TestAnagramMap(t *testing.T) {
	anagramMap := map[string][]string{
		"акптя":  {"пятак", "пятка", "тяпка"},
		"иклост": {"листок", "слиток", "столик"},
	}
	// Проверяем удаление множеств из одного элемента
	deleteSingleElementAnagrams(anagramMap)
	if len(anagramMap) != 1 {
		t.Errorf("Ожидалась длина карты 1, получено %d", len(anagramMap))
	}
}

func deleteSingleElementAnagrams(anagramMap map[string][]string) {
	for sortedWord, anagrams := range anagramMap {
		if len(anagrams) == 1 {
			delete(anagramMap, sortedWord)
		}
	}
}

func TestMain(m *testing.M) {
	// Проверка функции main не выводит ошибки при выполнении
	main()
}
