package handlers

import (
	config "bobby-website/server/configurations"
	fact "bobby-website/server/factory/admin"
	"html/template"
	"net/http"
)

//InitialRequest this is the first empty handler
func InitialRequest(wr http.ResponseWriter, req *http.Request) {
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
			//set the token in cookie
			cookie := newCookie("access_token", *token)
			http.SetCookie(wr, cookie)
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
		//set the token in cookie as empty
		cookie := newCookie("access_token", "")
		http.SetCookie(wr, cookie)
		wr.WriteHeader(200)
	}
}
