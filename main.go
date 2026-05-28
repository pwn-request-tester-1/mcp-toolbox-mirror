package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "MCP Toolbox Mirror - Conformance Test Server")
	})
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
