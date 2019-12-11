package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
	"unicode/utf8"
)

func handleGET(w http.ResponseWriter, r *http.Request) {
	filename := "index.html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}
func handlePOST(w http.ResponseWriter, r *http.Request) {
	log.Println("reader", r.FormValue("name"), r.FormValue("address"))
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	name := r.PostFormValue("name")
	address := r.PostFormValue("address")
	value := name + ":" + address
	if utf8.RuneCountInString(value) > 1 {
		log.Println(utf8.RuneCountInString(value))
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "token", Value: url.QueryEscape(value), Expires: expiration}
		http.SetCookie(w, &cookie)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleGET).Methods(http.MethodGet)
	r.HandleFunc("/", handlePOST).Methods(http.MethodPost)

	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
