package pattern

/*
	Реализовать паттерн «цепочка вызовов».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
type handler interface {
	execute(*item)
	setNext(handler)
}

type itemType int

const (
	Wear itemType = iota + 1
	Food
	Toy
)

// Будем по отдельности обрабатывать поля item в отдельных Handler-ах
type item struct {
	uid      string
	name     string
	itemType itemType
}

// Handler для поля uid структуры item
type uidHandler struct {
	uid  string
	next handler
}

func (uh *uidHandler) execute(item *item) {
	item.uid = uh.uid
	uh.next.execute(item)
}

func (uh *uidHandler) setNext(next handler) {
	uh.next = next
}

// Handler для поля name структуры item
type nameHandler struct {
	name string
	next handler
}

func (nh *nameHandler) execute(item *item) {
	item.name = nh.name
	nh.next.execute(item)
}

func (nh *nameHandler) setNext(next handler) {
	nh.next = next
}

// Handler для поля itemType структуры item
type itemTypeHandler struct {
	itemType itemType
	next     handler
}

func (r *itemTypeHandler) execute(item *item) {
	item.itemType = r.itemType
}

func (r *itemTypeHandler) setNext(next handler) {
	r.next = next
}

func ChainOfResp_main() {
	item := item{}
	itemTypeHandler := &itemTypeHandler{itemType: Toy}
	nameHandler := &nameHandler{name: "Toy - GiantHorse"}
	nameHandler.setNext(itemTypeHandler)
	uidHandler := &uidHandler{uid: "horsewb"}
	uidHandler.setNext(nameHandler)
	uidHandler.execute(&item)
}

//https://refactoring.guru/ru/design-patterns/chain-of-responsibility/go/example
