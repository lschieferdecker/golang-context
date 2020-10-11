package main

import (
	"fmt"
	"net/http"
)

func main() {
	finalHandler := http.HandlerFunc(testHandler)
	http.Handle("/test", addContextValueMiddleware(finalHandler))

	http.ListenAndServe(":8080", nil)
}

func testHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Context().Value("Name"))
}
