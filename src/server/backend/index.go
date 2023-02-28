package backend

import (
	"fmt"
	"net/http"
	"time"

	"2gte1.xyz/gcc/symparse/src/server/abs"
	"github.com/gin-gonic/gin"
	wcors "github.com/rs/cors/wrapper/gin"
)

/*the backend index page, interface for data input*/
type BackendPageServer struct {
}

func (bps *BackendPageServer) LaunchServer() {
	abs.EngGroup.Go(func() error {
		return (&http.Server{
			Addr:         ":" + abs.BackendPagePort,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler: func() http.Handler {
				eng := gin.New()
				eng.Use(gin.Recovery())

				//define cors rule
				eng.Use(wcors.New(wcors.Options{
					AllowedOrigins: []string{
						fmt.Sprintf("http://%s:%s", abs.MyIp, abs.HomePagePort),
						fmt.Sprintf("https://%s:%s", abs.MyIp, abs.HomePagePort),
						fmt.Sprintf("http://%s:%s", abs.MyDomain, abs.HomePagePort),
						fmt.Sprintf("https://%s:%s", abs.MyDomain, abs.HomePagePort),
					},
					AllowCredentials: true,
					Debug:            true,
				}))

				eng.GET("/", func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK,
						gin.H{
							"code": http.StatusOK,
							"data": "welcom to backend page",
						})
				})
				return eng
			}(),
		}).ListenAndServe()
	})
}

func NewBackendPageServer() abs.IServer {
	return new(BackendPageServer)
}
