package main

import ("fmt"
        "net/http"
        "database/sql"

      	_ "github.com/go-sql-driver/mysql")



func home_page(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Welcome!")
}

func handleRequest(){
  http.HandleFunc("/",home_page)
  http.ListenAndServe(":8080", nil)
}
func main()  {
  handleRequest()
     db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
     if err != nil {
       panic(err)
     }
     defer db.Close()
     fmt.Println("Connection with MySQL")
}
