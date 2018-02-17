package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	if flag.NArg() > 1 {
		log.Fatal("at most path argument is needed")
	}

	dir := "."
	if flag.NArg() == 1 {
		dir = flag.Arg(0)
	}

	http.Handle("/", http.FileServer(http.Dir(dir)))

	log.Fatal(http.ListenAndServe(*addr, nil))
}
