package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// name := "Alesker2"
	// city := "EKB"
	// myMap := make(map[string]string, 0)
	// myMap["Name"] = "Alesker2"
	// myMap["City"] = "EKB"
	// http.ServeFile(w, r, "templates/main.html")
	user := User{Name: "Peter", City: "Mahachkala"}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		fmt.Println(w, err)
		return
	}
	if err = tmpl.Execute(w, name); err != nil {
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is about page")
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is contacts page")
}

func main() {
	//mian page
	//about page
	//contacts
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contacts", contactsHandler)
	log.Println("Server starting")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
