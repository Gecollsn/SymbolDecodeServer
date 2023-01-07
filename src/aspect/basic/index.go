package basic

import (
	"net/http"
	"path/filepath"
	"time"

	"2gte1.xyz/gcc/symparse/src/aspect/api"
	"2gte1.xyz/gcc/symparse/src/filcom"
	"github.com/gin-gonic/gin"
)

/*home index page, interface for user access*/

type HomePageServer struct {
}

func (hps *HomePageServer) loadTemplates(eng *gin.Engine) {
	eng.LoadHTMLGlob(filepath.Join(filcom.Dirname(), "templates/**/**/*.html"))
}

func (hps *HomePageServer) LaunchServer() {
	EngGroup.Go(func() error {
		return (&http.Server{
			Addr:         ":" + HomePagePort,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			Handler: func() http.Handler {
				eng := gin.New()
				eng.Use(gin.Recovery())
				hps.loadTemplates(eng)
				eng.GET("/", func(ctx *gin.Context) {
					ctx.HTML(http.StatusOK, "index.html", gin.H{})
				})
				return eng
			}(),
		}).ListenAndServe()
	})
}

func NewHomePageServer() api.IServer {
	return new(HomePageServer)
}
