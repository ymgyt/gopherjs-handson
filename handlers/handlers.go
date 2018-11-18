package handlers

import (
	"github.com/ymgyt/gopherjs-handson/common"
)

// New -
func New(env *common.Env) *Handlers {
	return &Handlers{
		Example:  &Example{Env: env},
		GopherJS: &GopherJS{Env: env},
		Static:   &Static{Env: env},
	}
}

// Handlers -
type Handlers struct {
	Example  *Example
	GopherJS *GopherJS
	Static   *Static
}
