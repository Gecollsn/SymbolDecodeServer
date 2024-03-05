package main

import (
	"flag"
	"fmt"

	"gecollsn.it/symparse/server"
)

func main() {
	var isDebug = flag.Bool("debug", false, "Is current env in debug mode")

	flag.Parse()

	if *isDebug {
		fmt.Println("In debug mode!")
	}

	server.SymparseRun()
}
