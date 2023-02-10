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

}
