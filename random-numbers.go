package main

import (
	"net/http"
	"io"
	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		sessionID := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sessionID.String(),
		}
		http.SetCookie((w, c)
	}
	io.WriteString(w, c.String())
}

func bar (w http.ResponseWriter, req *http.Request){
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	c.Value = "something goes here"
	// http.SetCookie(w, c)
	io.WriteString(w, c.String())
}