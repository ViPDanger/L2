package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type flags struct {
	f string
	d string
	s bool
}

func main() {
	flags := FlagParser()
	if flags.f != "" {
		_, err := strconv.Atoi(flags.f)
		if err != nil {
			log.Fatalln("Can't read the field f")
		}
	}
	var stringLines []string
	wg := &sync.WaitGroup{}

	fmt.Println("Введите строки")
	wg.Add(1)
	go func() {
		for {
			var newLine string
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newLine = scanner.Text()
			if newLine != "" {
				if newLine == "/p" {
					wg.Done()
					break
				}

				stringLines = append(stringLines, newLine)
			}

		}
	}()

	time.Sleep(1 * time.Millisecond)
	wg.Wait()
	lines := Cut(flags, stringLines)
	for _, line := range lines {
		fmt.Println(line)
	}

}

func FlagParser() flags {
	var f = flag.String("f", "", "")
	var d = flag.String("d", "", "")
	var s = flag.Bool("s", false, "")
	flag.Parse()
	return flags{f: *f, d: *d, s: *s}
}

func Cut(flags flags, stringLines []string) []string {
	lines := []string{}
	separatedLines := [][]string{}
	delimiter := "	"
	field := 0

	if flags.d != "" {
		delimiter = flags.d
	}

	if flags.f != "" {
		f, err := strconv.Atoi(flags.f)
		if err != nil {
			log.Println("Can't read the field")
		}
		field = f
	}

	for _, val := range stringLines {
		if (!flags.s) || strings.Contains(val, delimiter) {
			separatedLines = append(separatedLines, strings.Split(val, delimiter))
		}
	}

	for i := 0; i < len(separatedLines); i++ {
		if flags.f == "" {
			lines = append(lines, separatedLines[i]...)
		} else {
			if len(separatedLines[i]) >= field {
				lines = append(lines, separatedLines[i][max(0, field-1)])
			}
		}
	}
	return lines
}
