package model

import "github.com/samber/mo"

type User struct {
	id       mo.Option[uint]
	githubID uint
	name     string
}

func (u *User) IsIDPresent() bool {
	return u.id.IsPresent()
}

func (u *User) GithubID() uint {
	return u.githubID
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

func (u *User) SetName(name string) {
	u.name = name
}

func NewUserWithoutID(githubID uint, name string) *User {
	return &User{
		githubID: githubID,
		name:     name,
	}
}

func NewUser(id uint, githubID uint, name string) *User {
	return &User{
		id:       mo.Some(id),
		githubID: githubID,
		name:     name,
	}
}
