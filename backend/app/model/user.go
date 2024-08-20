package model

import "github.com/samber/mo"

type User struct {
	id   mo.Option[uint]
	name string
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) ID() uint {
	if u.id.IsAbsent() {
		// None になるのはユーザーの新規作成時のみで, 基本起こり得ない
		// とはいえ, panic しないのが望ましい
		panic("ID is not set")
	}

	return u.id.MustGet()
}

func (u *User) Name() string {
	return u.name
}

func NewUserByName(name string) *User {
	return &User{name: name}
}

func NewUser(id uint, name string) *User {
	return &User{
		id:   mo.Some(id),
		name: name,
	}
}
