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
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error: %v", r)
			}
		}()

		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header

		if ctxHeader := r.Header.Get("X-Context"); ctxHeader == "" {

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Missing Header X-Context")
		} else {
			ctx := myContext{h: ctxHeader}
			result := h.handleRequest(&ctx)

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%v", result)
		}

	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
