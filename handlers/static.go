package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ymgyt/gopherjs-handson/common"
)

// Static -
type Static struct {
	fs http.Handler
	*common.Env
}

// StaticRoot -
func (s *Static) StaticRoot(root string, prefix string) *Static {
	s.fs = http.StripPrefix(prefix, http.FileServer(http.Dir(root)))

	return s
}

// Static -
func (s *Static) Static(c *gin.Context) {
	s.fs.ServeHTTP(c.Writer, c.Request)
}
