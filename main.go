package main

import (
	"flag"
	"fmt"

	"2gte1.xyz/gcc/symparse/src/server"
)

func env() {
	var isDebug = flag.Bool("debug", false, "Is current env in debug mode")

	flag.Parse()

	if *isDebug {
		fmt.Println("In debug mode!")
	}
}

func main() {
	env()
	server.StartServer()
}
