package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

  w.Write([]byte("Hello git hub :D"))
}

func main() {
  rtr := mux.NewRouter()
  rtr.HandleFunc("/user/{name:[a-z]+}/profile", profile).Methods("GET")

	rtr.HandleFunc("/hello", hello)

  http.Handle("/", rtr)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  name := params["name"]
  w.Write([]byte("Hello " + name))
}
