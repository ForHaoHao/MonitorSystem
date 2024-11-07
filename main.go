package main

import (
	"fmt"
	"net/http"
)
import "MonitorSystem/routers"

func main() {
	routers := routers.NewRouter()
	err := http.ListenAndServe(":8080", routers)

	if err != nil {
		fmt.Printf(err.Error())
	}
}
