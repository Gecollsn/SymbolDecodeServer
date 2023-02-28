package frontend

import (
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"2gte1.xyz/gcc/symparse/src/env"
	"2gte1.xyz/gcc/symparse/src/server/abs"
	"2gte1.xyz/gcc/symparse/src/tools/pathabt"
	"github.com/gin-gonic/gin"
)

/*home index page, interface for user access*/

type HomePageServer struct {
}

func (hps *HomePageServer) loadTemplates(eng *gin.Engine) {
	eng.LoadHTMLGlob(filepath.Join(pathabt.Dirname(), "templates/**/*.html"))
	eng.Static("/assets", filepath.Join(pathabt.Dirname(), "templates/home/assets"))
}

func (hps *HomePageServer) LaunchServer() {
	cfg := env.FrontCfg()

	abs.EngGroup.Go(func() error {
		return (&http.Server{
			Addr:         ":" + strconv.Itoa(cfg.Port),
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

func NewHomePageServer() abs.IServer {
	return new(HomePageServer)
}
