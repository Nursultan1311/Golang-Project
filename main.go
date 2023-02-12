package main

import (
  "database/sql"

  "fmt"

  "net/http"

  "html/template"

  _ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template

type User struct {
  id       int
  fullname string
  email    string
  phone    string
}

// When we run the main.go file firstly we call main method
func main() {
  tpl, _ = template.ParseGlob("templates/*.html")
  handleRequest() // we call another method named "handleRequest"
}

// this method checks all requests from the website
func handleRequest() {
  http.HandleFunc("/add_user", addUser)      // if domain name will extend like "localhost:8080/home_page" then we call homePage method
  http.ListenAndServe("localhost:8080", nil) //  this is main domain name
}

// this method has two parameters. ResponseWriter for adding html file to http, and second one for getting a request
func addUser(w http.ResponseWriter, r *http.Request) {
    // if r.Method == "POST" {
    tpl.ExecuteTemplate(w, "add_user.html", nil) // here we are writing add_user.html file with by Template
    // }

    r.ParseForm()
    surname := r.FormValue("surname")
    name := r.FormValue("name")
    email := r.FormValue("email")
    password := r.FormValue("password")

    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
    if err != nil {
      panic(err)
    }
    defer db.Close()

    insert, err := db.Query(fmt.Sprintf("INSERT INTO `users`(`name`, `surname`, `email`, `password`) VALUES ('%s', '%s', '%s', '%s')", name, surname, email, password))
    if err != nil {
      panic(err.Error())
    }
    defer insert.Close()

    fmt.Println("User was succesfully inserted")

    return

}
