package abs

import "golang.org/x/sync/errgroup"

var (
	EngGroup        errgroup.Group
	HomePagePort    = "8080"
	BackendPagePort = "8081"
	MyDomain        = "localhost"
	MyIp            = "127.0.0.1"
)

type IServer interface {
	LaunchServer()
}
