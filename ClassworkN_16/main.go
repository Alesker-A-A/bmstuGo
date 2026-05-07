package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	// http.NotFound(w, r)
	// 	// w.WriteHeader(http.StatusNotFound) //404
	// 	// http.Redirect(w, r, "/", http.StatusFound)
	// 	return
	// }
	http.ServeFile(w, r, "form.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	action := params["action"]
	city := params["city"]

	if action == "phone" && city == "msk" {
		w.Write([]byte("Телефон в Москве: 8(485)0000000"))
		return
	} else if action == "adress" && city == "spb" {
		w.Write([]byte("Адрес в Питере: Невский"))
		return
	} else if action == "adress" && city == "msk" {
		w.Write([]byte("Адрес в Москве: Первомайская"))
		return
	} else {
		w.Write([]byte("No such action, sorry"))
	}
}

func getHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my friend, I dont know ur email"))
}

func postHelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Hello %s, ur email: %s", name, email)
}

func main() {
	//gorilla/mux
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/about", aboutHandler)
	// /contacts/{phone/adress/}/{msk/spb}
	r.HandleFunc("/contacts/{action}/{city}", contactsHandler)
	r.HandleFunc("/hello", getHelloHandler).Methods("GET")
	r.HandleFunc("/hello", postHelloHandler).Methods("POST")

	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
