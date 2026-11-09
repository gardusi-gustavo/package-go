package main

import (
	"encoding/json"
	"log"
	"net/http"

	packagego "github.com/gardusi-gustavo/package-go"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]int{"answer": packagego.Answer()})
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
