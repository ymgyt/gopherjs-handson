package handlers

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ymgyt/gopherjs-handson/shared/model"

	"github.com/gin-gonic/gin"
	"github.com/ymgyt/gopherjs-handson/common"
	"go.isomorphicgo.org/go/isokit"
)

// Example -
type Example struct {
	*common.Env
}

// Example -
func (e *Example) Example(c *gin.Context) {
	e.TemplateSet.Render("example", &isokit.RenderParams{Writer: c.Writer, Data: nil})
}

func (e Example) ToLower(c *gin.Context) {
	const op = "example: ToLower"
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("%s failed to read request body %s\n", op, err)
		return
	}

	var text string
	if err = json.Unmarshal(b, &text); err != nil {
		fmt.Printf("%s failed to json unmarshal %s\n", op, err)
		return
	}

	res, err := json.Marshal(strings.ToLower(text))
	if err != nil {
		fmt.Printf("%s failed to json marshal %s", op, err)
		return
	}

	fmt.Printf("%s %s -> %s\n", op, b, res)
	c.Writer.Write(res)
}

// Gophers -
func (e *Example) Gophers(c *gin.Context) {

	fetchGopehrs := func() []*model.Gopher {

		gophers := []*model.Gopher{
			{Name: "A", Point: 100, Note: "isomorphic!"},
			{Name: "B", Point: 150, Note: "dom dom"},
			{Name: "C", Point: 80, Note: "keep doing !"},
			{Name: "D", Point: 200, Note: "golang or gohome"},
		}

		return gophers
	}

	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(fetchGopehrs()); err != nil {
		fmt.Printf("failed to gob encode %s\n", err)
		return
	}

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Write(b.Bytes())
}
