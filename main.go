package main

import (
	"log" // NOTE try slog when v1.21 is released
	"os"
)

func main() {
	// configure log
	f, err := os.Create("/tmp/toyls.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // this may return error so it is a little optimistic
	log.SetOutput(f)

	// initialize server
	log.Println("initializing server...")
}
