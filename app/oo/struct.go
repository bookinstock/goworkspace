package oo

import "fmt"

func RunStruct() {
	user := NewUser("u1", 18)
	fmt.Println("user=", user, user.name, user.age)
	user.setName("uuu")
	fmt.Println("name=", user.getName())
}

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	user := new(User)
	user.name = name
	user.age = age
	return user
}

func (u *User) getName() string {
	return u.name
}

func (u *User) setName(name string) {
	u.name = name
}
