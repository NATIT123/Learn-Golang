package logic

import "be-ep/internal/database"

type Hello struct {
	UserDataAccessor database.UserDataAccessor
}

func sayHello(thing string) (gretting string) {
	return "Hello" + thing
}

func SayHello(thing string) (gretting string) {
	return "Hello " + thing
}

func (h Hello) Hello() {}
