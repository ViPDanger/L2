Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	
	go func() {
		for {
			select {
			case v, ok := <-a:
				if ok{
				c <- v
				}
				
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
 Результат - случайный список значений с возрастанием чётных и нечётных значений (независимо), после чего softlock с бесконечным значением 0

 Пока каналы a и b находятся в гонке, значения из них будут выводиться. Однако когда один из каналов закрывается, чтение из него будет выводить default значение канала (0 в случае chan int) что будет приводить к чтению default значений из a и b в c.
 Для решения проблемы нужно добавить bool проверки чтения канала.


```
...

```
