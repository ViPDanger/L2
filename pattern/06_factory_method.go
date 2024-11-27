package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
import "fmt"

type product struct {
	uid  string
	name string
}

// Общее описание обьекта, создаваемого нами
type distribution struct {
	date string
	cost int
}

func (d *distribution) setDate(date string) {
	d.date = date
}
func (d *distribution) setCost(cost int) {
	d.cost = cost
}
func (d *distribution) getDate() string {
	return d.date
}
func (d *distribution) getCost() int {
	return d.cost
}

// Реализация обьекта
type aviaDistribution struct {
	distribution
}

func (d *aviaDistribution) distribute(p *product) {
	fmt.Println("Product with uid ", p.uid, " distibuted by air with cost ", d.cost)
}

func newAviaDistribution() distributionMethod {
	return &aviaDistribution{}
}

// Реализация обьекта
type railDistribution struct {
	distribution
}

func (d *railDistribution) distribute(p *product) {
	fmt.Println("Product with uid ", p.uid, " distibuted by rails with cost ", d.cost)
}
func newRailDistribution() distributionMethod {
	return &railDistribution{}
}

// Абстрактный обьект получаемый фабрикой
type distributionMethod interface {
	setDate(string)
	setCost(int)
	getDate() string
	getCost() int
	distribute(p *product)
}

// Фабрика
type distributioner struct {
}

func (d *distributioner) getDistributionMethod(i int) distributionMethod {
	switch i {
	case 0:
		return newAviaDistribution()
	default:
		return newRailDistribution()
	}
}

func Factory_main() {
	product := &product{uid: "123wb", name: "Toy"}
	// Создание фабрики
	distributioner := &distributioner{}
	// Создание обьектов фабрики
	distribution := distributioner.getDistributionMethod(0)
	// Использование обьекта
	distribution.setCost(10)
	distribution.distribute(product)
	// Новый обьект фабрики
	distribution = distributioner.getDistributionMethod(1)
	distribution.setCost(5)
	distribution.distribute(product)
}
