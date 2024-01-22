package server

import (
	"RedditAutoPost/config"
	"fmt"
	"net/http"
)

func StartServer(config config.Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Path: %s!", r.URL.Path[1:])
	})

	http.ListenAndServe(config.Server.Port, nil)
}
