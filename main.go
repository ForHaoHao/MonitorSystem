package main

import (
	"fmt"
	"net/http"

	"router"
)

func main() {
	routers := router.NewRouter()
	err := http.ListenAndServe(":8080", routers)

	if err != nil {
		fmt.Printf(err.Error())
	}
}
