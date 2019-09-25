package handlers

import (
	config "bobby-website/server/configurations"
	fact "bobby-website/server/factory/admin"
	"html/template"
	"net/http"
)

//NOTE: in all request we need to check if we have access token set in cookie.
//if access token is been cleared or been destroyed by closing the browser an call should be made
//in which we will set the access token in cookie an this can be handled using middle ware in routing

//Landing this is the first empty handler
func Landing(wr http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(config.HTMLFILEPATH + "/landing.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(wr, nil)
}

//https://thinkingeek.com/2018/05/31/setting-and-deleting-cookies-in-go/

//SignIn will be on login of admin
func SignIn(wr http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseForm()
		if token, err := fact.ValidateRequest(req.Form.Get("email"), req.Form.Get("password")); err == nil {
			//set the token in header to client
			wr.Write([]byte(*token))
			wr.WriteHeader(200)
		} else {
			wr.WriteHeader(401)
		}
	}
}

//SignOut will signout and will end current session
func SignOut(wr http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		fact.ADMIN.DestroyAccessToken()
		wr.WriteHeader(200)
	}
}

//CheckSession will act as middleware which will keep check if session token is removed from the cookie.
//if so it will call for new access token to backend
func CheckSession(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			//check access token which would be fetched from authorization header
			wr.Write([]byte(*fact.RefreshToken()))
			handler.ServeHTTP(wr, req)
			//will show us invalid token message
			wr.WriteHeader(401)
		}
	})
}
