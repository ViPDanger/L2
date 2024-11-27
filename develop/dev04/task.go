package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"sort"
	"strings"
)

func findAnagrams(arr *[]string) *map[string][]string {
	words := *arr
	for i, _ := range words {
		words[i] = strings.ToLower(words[i])
	}
	anagramsMap := make(map[string][]string)
	anagrams := []string{}
	for currentID := 0; currentID < len(words); currentID++ {

		var currentIsAnagram = false
		// Является ли текущее слово анаграммой
		for anagramaID := 0; anagramaID <= len(anagrams)-1; anagramaID++ {
			if isAnagram(words[currentID], anagrams[anagramaID]) {
				anagramsMap[anagrams[anagramaID]] = appendUnique(anagramsMap[anagrams[anagramaID]], words[currentID])
				currentIsAnagram = true
				break
			}
		}
		if !currentIsAnagram {
			anagramsMap[words[currentID]] = make([]string, 1)
			anagramsMap[words[currentID]][0] = words[currentID]
			anagrams = append(anagrams, words[currentID])
		}
	}

	for _, anagrama := range anagramsMap {
		sort.Strings(anagrama)
	}
	return &anagramsMap
}

func isAnagram(word1 string, word2 string) bool {
	intersections := map[rune]int{}
	for _, val := range word1 {
		intersections[val]++
	}
	for _, val := range word2 {
		intersections[val]++
	}
	for _, val := range intersections {
		if val < 2 {
			return false
		}
	}
	return true
}

func appendUnique(str []string, word string) []string {
	isUnique := true
	for _, r := range str {
		if r == word {
			isUnique = false
		}
	}
	if isUnique {
		str = append(str, word)
	}
	return str
}
