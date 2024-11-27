package pattern

import "log"

/*
	Реализовать паттерн «строитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Первичная структура user, которую будем получать из builder
type user struct {
	name         []byte
	password     []byte
	access_level int
}

// Интерфейс билдеров
type IUserBuilder interface {
	setPassword()
	setName()
	setAccessLevel()
	getUser() user
}

// User Builder
type UserBuilder struct {
	name         []byte
	password     []byte
	access_level int
}

func (builder *UserBuilder) setPassword() {
	builder.password = []byte("password")
}
func (builder *UserBuilder) setName() {
	builder.name = []byte("User")
}
func (builder *UserBuilder) setAccessLevel() {
	builder.access_level = 0
}
func (builder *UserBuilder) getUser() user {
	return user{
		name:         builder.name,
		password:     builder.password,
		access_level: builder.access_level,
	}
}

// Admin Builder
type AdminBuilder struct {
	name         []byte
	password     []byte
	access_level int
}

func (builder *AdminBuilder) setPassword() {
	builder.password = []byte("admin")
}
func (builder *AdminBuilder) setName() {
	builder.name = []byte("Admin")
}
func (builder *AdminBuilder) setAccessLevel() {
	builder.access_level = 1
}
func (builder *AdminBuilder) getUser() user {
	return user{
		name:         builder.name,
		password:     builder.password,
		access_level: builder.access_level,
	}
}

// Director
type Director struct {
	builder IUserBuilder
}

func (director *Director) SetBuilder(builder IUserBuilder) {
	director.builder = builder
}

func (director *Director) BuildUser() user {
	director.builder.setName()
	director.builder.setPassword()
	director.builder.setAccessLevel()
	return director.builder.getUser()
}

func Builder_main() {
	var director Director
	director.SetBuilder(&AdminBuilder{})
	log.Println(director.BuildUser())
	director.SetBuilder(&UserBuilder{})
	log.Println(director.BuildUser())
}

//
