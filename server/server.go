package server

import (
	config "bobby-website/server/configurations"
	h "bobby-website/server/handlers"
	"fmt"
	"net/http"
)

//Start will start the server
func Start() {
	fmt.Println("Server is listening on ", config.PORT)
	routers()
	http.ListenAndServe(config.PORT, nil)
}

func routers() {
	http.HandleFunc("/", h.InitialRequest)
}
