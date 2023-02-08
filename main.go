package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "Hello, World!")

		if err != nil {
			log.Println("Error:", err)
		}

		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	http.ListenAndServe(":8080", nil)
}
