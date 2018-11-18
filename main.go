package main

import (
	"fmt"
	"os"

	"github.com/ymgyt/gopherjs-handson/common"
	"github.com/ymgyt/gopherjs-handson/handlers"

	"github.com/gin-gonic/gin"
	"go.isomorphicgo.org/go/isokit"
)

const (
	defaultPort = "8123"
)

var (
	appRoot string
	port    string
)

func initializeTemplateSet(env *common.Env) {
	isokit.WebAppRoot = appRoot
	isokit.TemplateFilesPath = appRoot + "/shared/templates"
	isokit.StaticAssetsPath = appRoot + "/static"

	ts := isokit.NewTemplateSet()
	ts.GatherTemplates()

	env.TemplateSet = ts
}

func registerRoutes(env *common.Env, r *gin.Engine) {
	h := handlers.New(env)

	r.GET("/static/*filepath", h.Static.StaticRoot(appRoot+"/static", "/static").Static)
	r.GET("/js/client.js", h.GopherJS.AppRoot(appRoot).ClientJS)
	r.GET("/js/client.js.map", h.GopherJS.AppRoot(appRoot).ClientJSMap)
	r.GET("/example", h.Example.Example)
	r.GET("/example/gophers", h.Example.Gophers)
	r.POST("to-lower", h.Example.ToLower)
}

func main() {

	env := common.Env{}

	if appRoot == "" {
		fmt.Fprintln(os.Stderr, "the APP_ROOT environment variable required")
		os.Exit(1)
	}

	initializeTemplateSet(&env)

	r := gin.Default()
	registerRoutes(&env, r)

	r.Run(":" + port)
}

func init() {
	appRoot = os.Getenv("APP_ROOT")
	port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
}
