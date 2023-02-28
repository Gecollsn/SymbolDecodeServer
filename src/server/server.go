package server

import (
	"log"

	"2gte1.xyz/gcc/symparse/src/server/abs"
	"2gte1.xyz/gcc/symparse/src/server/backend"
	"2gte1.xyz/gcc/symparse/src/server/frontend"
)

func StartServer() {
	frontend.NewHomePageServer().LaunchServer()
	backend.NewBackendPageServer().LaunchServer()

	if err := abs.EngGroup.Wait(); err != nil {
		log.Fatal(err)
	}
}
