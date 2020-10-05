package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	handler := NewAppHandler("127.0.0.1:7722")

	log.Fatal(handler.Serve(":8080"))
}
