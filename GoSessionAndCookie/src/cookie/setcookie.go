package main

import (
	"net/http"
	"fmt"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {

	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Set Cookie Successful")

}

func getCookie(w http.ResponseWriter, r * http.Request) {
	cookie, _ := r.Cookie("username")
	fmt.Fprint(w, "GetCookieIs:", cookie)
}

func main() {
	http.HandleFunc("/getcookie", getCookie)
	http.HandleFunc("/setcookie", setCookie)
	err := http.ListenAndServe(":8888", nil)
	if(err != nil) {
		panic(err)
	}
}
