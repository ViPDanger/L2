package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
)

// Для применения паттерна фасада необходима струкура, к которой данный паттерн можно применить. Опишем её
// На примере некой бд клиентов, клиента и интерфейса вывода
type data struct {
	client
	data []byte
}

type database interface {
	Get(client client) []data
	Put(data data) error
	Delete(data data) error
}

type realdb struct {
	data []data
}

func (db *realdb) Get(client client) []data {
	data := make([]data, 0)
	for _, d := range db.data {
		if d.client.name == client.name {
			data = append(data, d)
		}
	}
	return data
}

func (db *realdb) Put(data data) error {
	db.data = append(db.data, data)
	return nil
}

func (db *realdb) Delete(data data) error {
	for i, _ := range db.data {
		if reflect.DeepEqual(data, db.data[i]) {
			db.data = append(db.data[:i], db.data[i+1:]...)
			return nil
		}
	}
	return errors.New("data not exists!")
}

type client struct {
	name     string
	password string
}

func (c *client) GetName() string {
	return c.name
}

func (c *client) CheckPassword(password string) bool {
	return password == c.password
}

func (c *client) NewPassword(OldPassword string, NewPassword string) error {
	if c.CheckPassword(OldPassword) {
		c.password = NewPassword
		return nil
	} else {
		return errors.New("Incorrect Old Passwod!")
	}
}

type Respond interface {
	io.ReadWriter
	Reset()
}

// Сам паттерн "фасад"

type Facade struct {
	DataBase database
	Respond  Respond
	Client   *client
}

func (f *Facade) New(DataBase database, Client *client, Respond Respond) {
	f.Client = Client
	f.DataBase = DataBase
	f.Respond = Respond
	f.Respond.Write([]byte("Created\n"))
}

func (f *Facade) Import(d []byte) {
	data := data{data: d}
	data.client = *f.Client
	f.Respond.Reset()
	err := f.DataBase.Put(data)
	if err != nil {
		f.Respond.Write([]byte(err.Error()))
	} else {
		f.Respond.Write([]byte("Import: "))
		f.Respond.Write(d)
	}
	fmt.Println(f.Respond)
}
func (f *Facade) Export() {
	f.Respond.Reset()
	f.Respond.Write([]byte("Export: \n"))
	for _, data := range (f.DataBase).Get(*f.Client) {
		f.Respond.Write(data.data)
		f.Respond.Write([]byte("\n"))
	}
	fmt.Println(f.Respond)
}
func (f *Facade) Delete(d []byte) {
	f.Respond.Reset()
	data := data{data: d}
	data.client = *f.Client
	err := f.DataBase.Delete(data)
	if err != nil {
		f.Respond.Write([]byte(err.Error()))
	} else {

		f.Respond.Write([]byte("Delete: "))
		f.Respond.Write(d)
	}
	fmt.Println(f.Respond)
}

// Применение фасада

func Facade_main() {
	b := make([]byte, 0)
	buff := bytes.NewBuffer(b)
	facade := Facade{}
	facade.New(&realdb{}, &client{name: "Пётр Петрович", password: "password"}, buff)
	facade.Import([]byte("New"))
	facade.Import([]byte("New2"))
	facade.Export()
	facade.Delete([]byte("New"))
	facade.Export()
}
