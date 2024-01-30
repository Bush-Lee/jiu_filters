package logic

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do your cors stuff
		// return if you do not want the FileServer handle a specific request
		enableCors(&w)
		// w.Header().Set("Content-Type", "application/json")

		fs.ServeHTTP(w, r)
	}
}

func StartServer(path string, port int) {
	fileServer := http.FileServer(http.Dir(path))
	http.Handle("/", cors(fileServer))
	portStr := strconv.Itoa(port)
	fmt.Println("Running server")
	log.Fatal(http.ListenAndServe(":"+portStr, nil))
}
