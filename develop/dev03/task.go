package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// //go run .\main.go -k3 -n -r -u .\sample.txt - пример команды
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	fileStrings := make(map[int]string)
	fileName := os.Args[len(os.Args)-1]
	// Определяем, существует ли файл
	_, err := os.Stat(fileName)
	if err != nil {
		log.Fatalf("File not found: %s", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	ScanFile(file, fileStrings)
	FlagParser(fileStrings)

	file.Close()
}

func ScanFile(file *os.File, fileStrings map[int]string) {
	fileScanner := bufio.NewScanner(file)

	i := 1
	for fileScanner.Scan() {
		fileStrings[i] = (fileScanner.Text())
		i++
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
}

func FlagParser(fileStrings map[int]string) {
	var k = flag.Int("k", 1, "k is ok!")
	var n = flag.Bool("n", false, "")
	var r = flag.Bool("r", false, "")
	var u = flag.Bool("u", false, "")
	flag.Parse()
	flagK(fileStrings, *k, *n, *r, *u)

}

func flagK(fileStrings map[int]string, k int, n, r, u bool) {
	if k < 1 {
		k = 1
	}
	fileStr := fileStrings
	// flag U
	if u {
		for i := 0; i < len(fileStr)-1; i++ {
			for j := i + 1; j < len(fileStr)-1; j++ {
				if fileStr[i] == fileStr[j] {
					delete(fileStr, j)
				}
			}
		}
	}
	//
	str := make([]string, len(fileStr))
	i := 0
	for _, v := range fileStr {
		str[i] = v
		i++
	}
	sort.Slice(str, func(i, j int) bool {
		re := regexp.MustCompile("[0-9]+")
		// flag N
		if n {
			var xi, xj int
			rei := re.FindAllString((str[i]), -1)
			rej := re.FindAllString((str[j]), -1)
			// flag K in N
			if k >= len(rei) {
				xi, _ = strconv.Atoi(rei[len(rei)-1])
			} else {
				xi, _ = strconv.Atoi(rei[k-1])
			}
			if k >= len(rej) {
				xj, _ = strconv.Atoi(rej[len(rej)-1])
			} else {
				xj, _ = strconv.Atoi(rej[k-1])
			}
			// flag r in N
			return (xi < xj) != r
		}
		// flag R,K in !N
		return ((str[i])[k-1:] < (str[j])[k-1:]) != r
	})
	fmt.Println(str)

	file, err := os.OpenFile(os.Args[len(os.Args)-1], os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	datawriter := bufio.NewWriter(file)
	for i, _ := range str {
		_, _ = datawriter.WriteString(str[i])
		if i < len(str)-1 {
			_, _ = datawriter.WriteString("\n")
		}
	}
	datawriter.Flush()
	file.Close()
}
