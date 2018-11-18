package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ymgyt/gopherjs-handson/common"
)

// GopherJS -
type GopherJS struct {
	appRoot string
	*common.Env
}

// AppRoot -
func (g *GopherJS) AppRoot(s string) *GopherJS {
	g.appRoot = s
	return g
}

// ClientJS -
func (g *GopherJS) ClientJS(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, g.appRoot+"/client/client.js")
}

// ClientJSMap -
func (g *GopherJS) ClientJSMap(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, g.appRoot+"/client/client.js")
}
