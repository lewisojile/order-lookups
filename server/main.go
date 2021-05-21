package main

import (
	"log"

	"github.com/lewisojile/order-app/boot"
)

func main() {

	err := boot.Start()

	if err != nil {
		log.Fatalln(err)
	}

}
