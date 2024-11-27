package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Strategy_pattern
*/
// Context, некий обьект пользующийся стратегией
type storage struct {
	products []product
	recieveStratagy
}

func (s *storage) store(p *product) {
	s.products = append(s.products, *s.recieve(p))
}

// Интерфейс stratagy
type recieveStratagy interface {
	recieve(p *product) *product
}

// Реализация stratagy
type recieveByAir struct {
}

func (r *recieveByAir) recieve(p *product) *product {
	fmt.Println("Product", p.uid, "recieved by air")
	return p
}

// Реализация stratagy
type recieveByRail struct {
}

func (r *recieveByRail) recieve(p *product) *product {
	fmt.Println("Product", p.uid, "recieved by rail")
	return p
}

func Strategy_main() {
	product := &product{uid: "wb123", name: "toy"}
	storage := storage{}
	storage.recieveStratagy = &recieveByAir{}
	storage.store(product)
	storage.recieveStratagy = &recieveByRail{}
	storage.store(product)
}
