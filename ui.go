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

func steps(w http.ResponseWriter, r *http.Request) {
	hdrs(w)
	w.Write([]byte(bingraph))
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

	rtr.PathPrefix("/d1/").Handler(http.StripPrefix("/d1/",
		http.FileServer(http.Dir("./converter1/"))))

//	rtr.HandleFunc("/b/f.json",

//	http.FileServer(http.Dir("./converter1/o.json"))

	rtr.PathPrefix("/v/").Handler(http.StripPrefix("/v/",
		http.FileServer(http.Dir("./vendor/"))))

	err := http.ListenAndServe(":3000", rtr)
	if err != nil {
		log.Fatal(err)
	}
}
