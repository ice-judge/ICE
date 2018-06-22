package router

import (
	"fmt"
	"net/http"
)

func ProblemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}

func HelloWorld() {
	fmt.Println("Hello, World!")
}
