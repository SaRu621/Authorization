package main

import (
	"Authorization/pkg/interfaces"
	Auth "Authorization/pkg/models/auth"
	"fmt"
)

func f(v interfaces.WriteDB) {
	fmt.Println("It works!", v)
}

func main() {
	t := Auth.Auth{}
	f(&t)
}
