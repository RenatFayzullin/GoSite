package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, happiness float64
	Hobbies               []string
}

func (user *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s, Age is %d, Money is %d", user.Name, user.Age, user.Money)
}

func (user *User) setNewName(name string) {
	user.Name = name
}

func main() {
	handlerRequest()
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Soccer", "Dance"}}
	bob.setNewName("Renat")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)

}

func handlerRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts")
}
