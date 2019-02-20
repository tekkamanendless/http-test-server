package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	portNumber := flag.Int("port", 8080, "Port number to listen on.")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		fmt.Printf("Path: %s\n", path)
		if strings.HasPrefix(path, "/status/") {
			path := strings.TrimPrefix(r.URL.Path, "/status/")
			path = strings.Trim(path, "/")
			statusCode, err := strconv.ParseInt(path, 10, 32)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(int(statusCode))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unhandled path"))
			return
		}
	})

	fmt.Printf("Listening on port %d.\n", *portNumber)
	http.ListenAndServe(fmt.Sprintf(":%d", *portNumber), nil)
}
