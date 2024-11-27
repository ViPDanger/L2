package pattern

import (
	"errors"
	"fmt"
	"log"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// State Interface
type state interface {
	stateAction(*productState) error
}

// Concrete States
type onStorage struct {
	storageID string
}

func (s *onStorage) stateAction(productState *productState) error {
	if s == nil || productState == nil {
		return errors.New("State: stateAction() nill pointer")
	}
	fmt.Println(productState.name, " on storage with StorageID ", s.storageID)
	return nil
}

type onDelivery struct {
	deliveryID string
}

func (s *onDelivery) stateAction(productState *productState) error {
	if s == nil || productState == nil {
		return errors.New("State: stateAction() nill pointer")
	}
	fmt.Println(productState.name, " on delivery with id ", s.deliveryID)
	return nil
}

type onPickUpPoint struct {
	pickUpPointID string
}

func (s *onPickUpPoint) stateAction(productState *productState) error {
	if s == nil || productState == nil {
		return errors.New("State: stateAction() nill pointer")
	}
	fmt.Println(productState.name, " on PickUpPoint with id ", s.pickUpPointID, " waiting for client")
	return nil
}

// State Context
type productState struct {
	product
	state state
}

func (s *productState) GetState() (state, error) {
	if s == nil {
		return nil, errors.New("Product State: GetState() nill pointer")
	}
	if s.state == nil {
		return nil, errors.New("Product State: GetState() state nil")
	}
	return s.state, nil
}
func (s *productState) SetState(newState state) error {
	if s == nil {
		return errors.New("Product State: SetState() nill pointer")
	}
	s.state = newState
	return nil
}

func (s *productState) applyStateAction() {
	state, err := s.GetState()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = state.stateAction(s)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func NewProductState(product product, state state) *productState {
	return &productState{
		product: product,
		state:   state,
	}
}

func State_main() {
	productState := NewProductState(product{uid: "123wb", name: "Horse Toy"}, &onStorage{storageID: "wbstorage1"})
	// Применение действия состояния
	productState.applyStateAction()
	// Изменение состояния
	err := productState.SetState(&onDelivery{deliveryID: "wbdelivery2"})
	if err != nil {
		log.Fatalln(err.Error())
	}
	// Применение действия состояния
	productState.applyStateAction()
	// Изменение состояния
	err = productState.SetState(&onPickUpPoint{pickUpPointID: "wbPickUpPoint3"})
	if err != nil {
		log.Fatalln(err.Error())
	}
	// Применение действия состояния
	productState.applyStateAction()

}
