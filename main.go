package main

import (
	"fmt"
	"net/http"

	"hitalent/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.Register(mux)

	addr := ":8080"
	fmt.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println(err)
	}
}
