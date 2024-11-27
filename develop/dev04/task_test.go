package main

import (
	"testing"
)

func TestFindAnagrams01(t *testing.T) {
	// Test 1
	arr := &[]string{"пятка", "Пятак", "тяпка", "листок", "слиток", "столик", "диван", "арбалет", "иванд", "тяпка", "диван", "Тяпка"}
	expected := map[string][]string{"арбалет": {"арбалет"}, "диван": {"диван", "иванд"}, "листок": {"листок", "слиток", "столик"}, "пятка": {"пятак", "пятка", "тяпка"}}
	defaultTest(t, arr, expected)
}

func TestFindAnagrams02(t *testing.T) {
	// Test 1
	arr := &[]string{"Затоп", "патоз", "Злоба", "базол", "платок", "плотка", "толпак", "робот", "ТОРОБ"}
	expected := map[string][]string{"затоп": {"затоп", "патоз"}, "злоба": {"базол", "злоба"}, "платок": {"платок", "плотка", "толпак"}, "робот": {"робот", "тороб"}}
	defaultTest(t, arr, expected)
}

// Функция проверки
func defaultTest(t *testing.T, arr *[]string, expected map[string][]string) {
	result := findAnagrams(arr)
	// проверка на expected <= result
	for anagrama, res := range *result {
		for i, _ := range res {
			_, ok := expected[anagrama]
			if !ok {
				t.Fatal("No ", anagrama, "anagrams in expected map")
			}
			if res[i] != expected[anagrama][i] {
				t.Fatal("Incorrect result. Expected", expected[anagrama][i], "; Result", res[i])
			}
		}
	}
	// проверка на result <= expected
	for anagrama, exp := range expected {
		for i, _ := range exp {
			_, ok := (*result)[anagrama]
			if !ok {
				t.Fatal("No ", anagrama, "anagrams in result map")
			}
			if exp[i] != (*result)[anagrama][i] {
				t.Fatal("Incorrect result. Expected", expected[anagrama][i], "; Result", exp[i])
			}
		}
	}
}
