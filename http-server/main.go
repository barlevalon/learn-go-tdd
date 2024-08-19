package main

import (
	"log"
	"net/http"
)


func main() {
  server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":6060", server))
}
