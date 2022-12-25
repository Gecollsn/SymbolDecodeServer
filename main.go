package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func env() {
	var isDebug = flag.Bool("debug", false, "whether current in debug mode")

	flag.Parse()

	if *isDebug {
		fmt.Println("In debug mode!")
	}
}

func serverRun() {
	engine := gin.Default()
	engine.Run()
}

func main() {
	env()
	serverRun()
}
