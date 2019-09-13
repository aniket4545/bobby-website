package main

import (
	config "bobby-website/server/configurations"
	h "bobby-website/server/handlers"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is listening on ", config.PORT)
	http.HandleFunc("/", h.InitialRequest)
	http.ListenAndServe(config.PORT, nil)
}
