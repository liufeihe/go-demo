package basic

import "net/http"

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})
	http.ListenAndServe(":8080", nil)
}
