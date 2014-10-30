package main

import (
    "flag"
	"github.com/gorilla/mux"
    "io"
    "log"
    "net"
    "net/http"
    "net/http/fcgi"
    "runtime"
)

func iiii(i interface {}) {
i=flag.Flag{};i=net.IP{};fcgi.Serve(nil,nil)
}

func hdrs(w http.ResponseWriter) {
    headers := w.Header()
    headers.Add("Content-Type", "text/html")
}

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
}

func binaries(w http.ResponseWriter, r *http.Request) {
	hdrs(w)
	w.Write([]byte(bingraph))
}

func home(w http.ResponseWriter, r *http.Request) {
	hdrs(w)
	io.WriteString(w, "<html><head></head><body><p>It works!</p></body></html>")
}

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", home)
	rtr.HandleFunc("/b", binaries)
	rtr.HandleFunc("/b/f.json", binaries)

	rtr.PathPrefix("/v/").Handler(http.StripPrefix("/v/",
		http.FileServer(http.Dir("./vendor/"))))

	err := http.ListenAndServe(":3000", rtr)
	if err != nil {
		log.Fatal(err)
	}
}

/*
package main

import (

)



func main() {
  rtr := mux.NewRouter()

	rtr.HandleFunc("/b", binaries)
//
  http.Handle("/", rtr)


  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}
*/
