Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
	результат выведения программы - error
	Причина схожа с listing02 - в переменную интерфейса мы передеём в любом случае реализующую его ссылку на структуру customError. (func test() возвращает переменную типа *customError)
	Для того, чтобы программа выводила значения корректно, вывод фунции должен быть интерфейсом error, а выходяшее из него значение nil/*customError
...

```
