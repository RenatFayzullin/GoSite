package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

func (user *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s, Age is %d, Money is %d", user.name, user.age, user.money)
}

func (user *User) setNewName(name string) {
	user.name = name
}

func main() {
	handlerRequest()
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8}
	bob.setNewName("Renat")
	fmt.Fprintf(w, bob.getAllInfo())
	fmt.Fprintf(w, "Hello")
	//tmpl, _ := teamplate.parseFiles()

}

func handlerRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts")
}
