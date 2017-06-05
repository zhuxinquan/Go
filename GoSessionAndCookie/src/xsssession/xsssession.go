package main

import (
	"net/http"
	"github.com/astaxie/session"
	"html/template"
)

var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}
func main() {
	http.HandleFunc("/", count)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}
