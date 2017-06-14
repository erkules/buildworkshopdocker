package main

import (
  "fmt"
  "net/http"
  //"strings"
  "log"
)

func sayHallo(w http.ResponseWriter, r *http.Request){
  r.ParseForm()
  fmt.Println(r.Form)
  //fmt.Println("path",r.URL.PATH)
  fmt.Fprintf(w,"Hello Dockerbuild \\o/")
}

func main(){
  http.HandleFunc("/",sayHallo)
  err := http.ListenAndServe(":80",nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
