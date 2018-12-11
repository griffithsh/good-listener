package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

var (
	portFlag = flag.String("addr", ":8080", "provide a port to listen on")
)

var handle http.HandlerFunc = func(w http.ResponseWriter, req *http.Request) {
	b, err := httputil.DumpRequest(req, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("DumpRequest: %v", err)))
		return
	}
	fmt.Println(string(b))

	w.WriteHeader(200)
}

func main() {
	flag.Parse()

	err := http.ListenAndServe(*portFlag, handle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "listenAndServe: %v\n", err)
	}
}
