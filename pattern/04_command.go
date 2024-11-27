package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import "fmt"

// Интерфейс Command для структуры button
type command interface {
	Call()
}

type button struct {
	command command
}

// Вызов Call() для команды в кнопке
func (b *button) press() {
	b.command.Call()
}

type Command struct {
	list *ShopList
	uid  string
}

func (c *Command) Call() {
	c.list.add(c.uid)
}

// Будем добавлять в shoplist предметы по нажатию кнопки
type ShopList struct {
	ItemList []string
}

func (sl *ShopList) add(uid string) {
	sl.ItemList = append(sl.ItemList, uid)
	fmt.Println("Item with uid ", uid, " added")
}

// Применение command паттерна
func Command_main() {
	shoplist := &ShopList{}

	command := &Command{
		list: shoplist,
		uid:  "123wb",
	}

	Button := &button{
		command: command,
	}
	Button.press()
}
