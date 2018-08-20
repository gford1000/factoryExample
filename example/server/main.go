package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Wire up application
	h := handler{}
	h.initialise()

	handler := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error: %v", r)
			}
		}()

		if r.Header == nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Missing header X-Context")

		} else {
			ctx := myContext{h: r.Header.Get("X-Context")}
			result := h.handleRequest(&ctx)

			w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%v", result)
		}
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
