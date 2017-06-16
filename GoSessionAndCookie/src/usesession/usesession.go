package main

import (
	"net/http"
	"github.com/astaxie/session"
	"fmt"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory","gosessionid",3600)
}

func getSession(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	session := sess.Get("username")
	fmt.Fprintf(w, "Get Session Ok, Session is %s", session)
}

func index(w http.ResponseWriter, r * http.Request) {
	sess := globalSessions.SessionStart(w, r)
	sess.Set("username", "set username")
	fmt.Fprint(w, "Set Session Ok!")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/get", getSession)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
