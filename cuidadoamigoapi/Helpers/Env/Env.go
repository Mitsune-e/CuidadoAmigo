package Env

import (
	"fmt"

	"github.com/gofor-little/env"
)

//path =: "../../.env"

func CheckEnv() {
	loadrr := env.Load(("settings.env"))
	if loadrr != nil {
		panic(loadrr)
	}
	HOST, err := env.MustGet("DB_HOST")
	if err != nil {
		panic(err)
	}
	fmt.Println(HOST)
}
