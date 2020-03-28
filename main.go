package main

import (
	"go-elastic-hotels/server"
	"log"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
