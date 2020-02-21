package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Package struct {
	V string `json:"v"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping")
}

func ping2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, ok := r.URL.Query()["id"]
	if !ok || len(id) == 0 {
		fmt.Fprintf(w, "Error")
	}
	p := Package{id[0]}
	js, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(js)
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/ping2", ping2)
	http.ListenAndServe(":8080", nil)
}
