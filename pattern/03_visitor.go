package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

import (
	"log"
)

type Item interface {
	getType() string
	accept(visitor)
}

// food struct
type food struct {
	name string
}

func (f *food) getType() string {
	return "Food"
}
func (f *food) accept(v visitor) {
	v.visitForFood(f)

}

// wear struct
type wear struct {
	name string
	size string
}

func (w *wear) getType() string {
	return "Wear"
}
func (w *wear) accept(v visitor) {
	v.visitForWear(w)
}

// visitor
type visitor interface {
	visitForFood(*food)
	visitForWear(*wear)
}

type nameVisitor struct {
	Name string
}

func (v *nameVisitor) visitForFood(f *food) {
	v.Name = f.name
}

func (v *nameVisitor) visitForWear(w *wear) {
	v.Name = w.name + "; size: " + w.size
}

func Visitor_main() {
	Visitor := &nameVisitor{}
	apple := &food{name: "Apple"}
	hoodie := &wear{name: "Hoodie", size: "XL"}
	apple.accept(Visitor)
	log.Println(Visitor.Name)
	hoodie.accept(Visitor)
	log.Println(Visitor.Name)

}
