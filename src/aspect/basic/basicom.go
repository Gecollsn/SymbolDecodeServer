package basic

import (
	"log"

	"golang.org/x/sync/errgroup"
)

var (
	EngGroup        errgroup.Group
	HomePagePort    = "8080"
	BackendPagePort = "8081"
	MyDomain        = "localhost"
	MyIp            = "127.0.0.1"
)

func StartServer() {
	NewHomePageServer().LaunchServer()
	NewBackendPageServer().LaunchServer()

	if err := EngGroup.Wait(); err != nil {
		log.Fatal(err)
	}
}
