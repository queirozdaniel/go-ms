package main

import (
	"upse/authentication/internal/config"
)

func main() {
	_, err := config.InitConfig(".")
	if err != nil {
		panic(err)
	}

}
