package main

import (
	router "./router"
)

func main() {
	r := router.InitRouter()
	r.Run()
}
