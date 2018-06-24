package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:3000", "the address")
	flag.Parse()

	distDir := "./dist"
	if len(os.Args) == 2 {
		distDir = os.Args[1]
	}

	if _, err := os.Stat(distDir); os.IsNotExist(err) {
		log.Fatalf("%s doesn't exist", distDir)
	}

	http.Handle("/", http.FileServer(http.Dir(distDir)))
	log.Printf("Serving %s", distDir)
	log.Printf("Listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
