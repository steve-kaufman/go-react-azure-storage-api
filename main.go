package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func signature(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["bloburi"]
	if !ok || len(keys[0]) < 1 {
		log.Println("'bloburi' is missing")
		return
	}

	bloburi := string(keys[0])

	token := GetToken(bloburi, os.Getenv("SAS_PASSWORD"))

	response := bloburi + "?" + token

	log.Println("SAS URI: " + response)

	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/signature", signature)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
