package main

import "testing"

func TestUnpack01(t *testing.T) {
	// Test 1
	str := "a4bc2d5e"
	expected := "aaaabccddddde"
	result, err := Unpack(str)
	if result != expected || err != nil {
		t.Fatal("Incorrect result. Expected", expected, "; Result", result)
	}

}

// Test 2
func TestUnpack02(t *testing.T) {

	str := "abcd"
	expected := "abcd"
	result, err := Unpack(str)
	if result != expected || err != nil {
		t.Fatal("Incorrect result. Expected", expected, "; Result", result)
	}
}

func TestUnpack03(t *testing.T) {

	str := "45"
	expected := ""
	result, err := Unpack(str)
	if err == nil || result != expected {
		t.Fatal("Incorrect result. Error == nil")
	}
}

func TestUnpack04(t *testing.T) {
	str := ""
	expected := ""
	result, err := Unpack(str)
	if result != expected || err != nil {
		t.Fatal("Incorrect result. Expected", expected, "; Result", result)
	}
}

func TestUnpack05(t *testing.T) {
	str := "qwe\\4\\5"
	expected := "qwe45"
	result, err := Unpack(str)
	if result != expected || err != nil {
		t.Fatal("Incorrect result. Expected", expected, "; Result", result)
	}
}
func TestUnpack06(t *testing.T) {
	str := "qwe\\45"
	expected := "qwe44444"
	result, err := Unpack(str)
	if result != expected || err != nil {
		t.Fatal("Incorrect result. Expected", expected, "; Result", result)
	}
}

func TestUnpack07(t *testing.T) {
	str := "qwe\\\\5"
	expected := "qwe\\\\\\\\\\"
	result, err := Unpack(str)
	if result != expected || err != nil {
		t.Fatal("Incorrect result. Expected", expected, "; Result", result)
	}
}

/*
	=== Задача на распаковку ===

	Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

		- "abcd" => "abcd"
		- "45" => "" (некорректная строка)
		- "" => ""
	Дополнительное задание: поддержка escape - последовательностей
		- qwe\4\5 => qwe45 (*)
		- qwe\45 => qwe44444 (*)
		- qwe\\5 => qwe\\\\\ (*)

	В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

	Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
