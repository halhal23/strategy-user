package main

import (
	"strategy-user/infrastructure"
)

func main() {
	r := infrastructure.NewRouter()
	r.Run()
}