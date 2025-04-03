Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Ответ test() output 2
	  anotherTest() output 1


	  В случае test() изменяемое поле в defer - return функции
	  Тогда как в anotherTest() изменяется внутренняя переменная функции, тогда как значение передалось в return до выполнения defer.
...

```
