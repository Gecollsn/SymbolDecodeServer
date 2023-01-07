package main

import (
	"flag"
	"fmt"

	"2gte1.xyz/gcc/symparse/src/aspect/basic"
)

func env() {
	var isDebug = flag.Bool("debug", false, "whether current in debug mode")

	flag.Parse()

	if *isDebug {
		fmt.Println("In debug mode!")
	}
}

func main() {
	env()
	basic.StartServer()
}
