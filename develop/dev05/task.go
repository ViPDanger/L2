package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type flags struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

func main() {
	fileName := os.Args[len(os.Args)-1]
	_, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("file isnot found", err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	Search(ScanFile(file), FlagParser())

	file.Close()
}

func ScanFile(file *os.File) []string {
	fileScanner := bufio.NewScanner(file)
	fileStrings := make([]string, 0)
	for fileScanner.Scan() {
		fileStrings = append(fileStrings, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	return fileStrings
}

func FlagParser() flags {
	var A = flag.Int("A", 0, "")
	var B = flag.Int("B", 0, "")
	var C = flag.Int("C", 0, "")
	var c = flag.Bool("c", false, "")
	var i = flag.Bool("i", false, "")
	var v = flag.Bool("v", false, "")
	var F = flag.Bool("F", false, "")
	var n = flag.Bool("n", false, "")
	flag.Parse()
	flags := flags{A: *A, B: *B, C: *C, c: *c, i: *i, v: *v, F: *F, n: *n}
	return flags
}

type row struct {
	id       int
	Row      string
	Included bool
}

func (r *row) include() {
	r.Included = true
}

func Search(fileStrings []string, flags flags) {
	str := os.Args[len(os.Args)-2]
	fs := make([]row, 0)
	for i, s := range fileStrings {
		fs = append(fs, row{id: i, Row: flagI(flags.i, s), Included: false})
	}

	// flag -i "ignore-case"

	// flag -n add row number

	for i := 0; i < len(fs); i++ {
		if flagF(flags.F, fs[i].Row, str) != flags.v {
			for j := max(0, min(i-flags.A, i-flags.C)); j <= min(len(fs)-1, max(i+flags.B, i+flags.C)); j++ {
				fs[j].include()
			}
		}
	}
	counter := 0
	for _, val := range fs {
		if val.Included {
			counter++
			// flag -n add row number
			if flags.n {
				fmt.Printf("row #%2d)  ", val.id)
			}
			fmt.Printf("%s\n", val.Row)
		}
	}
	if flags.c {
		fmt.Printf("Total rows count: %d\n", counter)
	}

}

// flag -i - "ignore-case" (игнорировать регистр)
func flagI(flag bool, str string) string {
	if flag {
		return strings.ToLower(str)
	} else {
		return str
	}
}

// flag -F - "fixed", точное совпадение со строкой, не паттерн
func flagF(flag bool, str, substr string) bool {
	if flag {
		for _, r := range strings.Fields(str) {
			if r == substr {
				return true
			}
		}
		return false
	}
	return strings.Contains(str, substr)
}
