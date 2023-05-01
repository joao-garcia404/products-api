package main

import "github.com/joao-garcia404/products-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")

	println(config.DBUser)
}
