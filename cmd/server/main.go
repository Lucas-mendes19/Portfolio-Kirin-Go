package main

import (
	"fmt"
	"github.com/Lucas-mendes19/Portfolio-Kirin-Go/configs"
)

func main() {
	env, _ := configs.Load(".")

	fmt.Println(env.ServerPort)
	fmt.Println(env.Database.Port)
}
