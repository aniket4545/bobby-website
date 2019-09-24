package server

import (
	config "bobby-website/server/configurations"
	h "bobby-website/server/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Start will start the server
func Start() {
	fmt.Println("Server is listening on ", config.PORT)
	http.Handle("/", routers())
	http.ListenAndServe(config.PORT, nil)
}

func routers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", h.InitialRequest)
	router.HandleFunc("/{file}", h.HandleFile)

	//all the rest api below would be gone through the session check
	router.HandleFunc("/signin", h.SignIn)
	router.HandleFunc("/signout", h.SignOut)

	router.Use(h.CheckSession)
	return router
}
