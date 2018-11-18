package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/ymgyt/gopherjs-handson/shared/model"

	"honnef.co/go/js/xhr"

	"github.com/gopherjs/gopherjs/js"

	"honnef.co/go/js/dom"
)

func run() {

	println("start client !")

	d := dom.GetWindow().Document()

	toggleCSSProperty(d)
	alertEvent(d)
	xhrLowercase(d)
	renderTable(d)
	localStorage(d)
}

func toggleCSSProperty(d dom.Document) {

	toggle := func(d dom.Document, show bool) {
		imageEle := d.GetElementByID("isomorphicGopher").(*dom.HTMLImageElement)

		var display string
		if show {
			display = "inline"
		} else {
			display = "none"
		}

		imageEle.Style().SetProperty("display", display, "")
	}

	show := d.GetElementByID("showGopher").(*dom.HTMLButtonElement)
	show.AddEventListener("click", false, func(_ dom.Event) {
		toggle(d, true)
	})

	hide := d.GetElementByID("hideGopher").(*dom.HTMLButtonElement)
	hide.AddEventListener("click", false, func(_ dom.Event) {
		toggle(d, false)
	})
}

func alertEvent(d dom.Document) {
	b := d.GetElementByID("alertMessageButton").(*dom.HTMLButtonElement)
	b.AddEventListener("click", false, func(e dom.Event) {
		v := d.GetElementByID("messageInput").(*dom.HTMLInputElement).Value
		js.Global.Call("alert", v)
	})
}

func xhrLowercase(d dom.Document) {

	lowercaseTextTransform := func() {
		input := d.GetElementByID("textToLowercase").(*dom.HTMLInputElement)

		b, err := json.Marshal(input.Value)
		if err != nil {
			fmt.Printf("failed to json.Marshal input value %s\n", err)
			return
		}

		data, err := xhr.Send("POST", "/to-lower", b)
		if err != nil {
			fmt.Printf("failed to xhr.Send %s\n", err)
			return
		}

		var s string
		if err = json.Unmarshal(data, &s); err != nil {
			fmt.Printf("failed to json.Unmarshal response data %s\n", err)
			return
		}

		input.Set("value", s)
	}

	b := d.GetElementByID("lowercaseTransformButton").(*dom.HTMLButtonElement)
	b.AddEventListener("click", false, func(_ dom.Event) {
		go lowercaseTextTransform()
	})
}

func renderTable(d dom.Document) {
	println("render table")

	fetchGophers := func() []*model.Gopher {
		data, err := xhr.Send("GET", "/example/gophers", nil)
		if err != nil {
			fmt.Printf("faield to fetch gophers %s\n", err)
			return nil
		}

		var gophers []*model.Gopher
		if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&gophers); err != nil {
			fmt.Printf("failed to decode gophers %s\n", err)
			return nil
		}

		return gophers
	}

	const tmplSrc = `
		<td>{{.Name}}</td>
		<td>{{.Point}}</td>
		<td>{{.Note}}</td>
	`

	tmpl, err := template.New("model").Parse(tmplSrc)
	if err != nil {
		fmt.Printf("failed to parse template %s\n", err)
		return
	}

	gophers := fetchGophers()

	tb := d.GetElementByID("tableBody")
	for i := 0; i < len(gophers); i++ {
		var b bytes.Buffer
		tmpl.Execute(&b, gophers[i])

		tr := d.CreateElement("tr")
		tr.SetInnerHTML(b.String())
		tb.AppendChild(tr)
	}
}

func localStorage(d dom.Document) {

	ls := js.Global.Get("localStorage")

	display := func() {
		list := d.GetElementByID("lsItemList")
		list.SetInnerHTML("")

		for i := 0; i < ls.Length(); i++ {
			key := ls.Call("key", i)
			value := ls.Call("getItem", key)

			dtEle := d.CreateElement("dt")
			dtEle.SetInnerHTML(key.String())

			ddEle := d.CreateElement("dd")
			ddEle.SetInnerHTML(value.String())

			list.AppendChild(dtEle)
			list.AppendChild(ddEle)
		}
	}

	save := func() {
		key := d.GetElementByID("lsItemKey").(*dom.HTMLInputElement)
		value := d.GetElementByID("lsItemValue").(*dom.HTMLInputElement)

		if key.Value == "" {
			return
		}

		ls.Call("setItem", key.Value, value.Value)

		key.Value = ""
		value.Value = ""

		display()
	}

	clear := func() {
		ls.Call("clear")
		display()
	}

	saveBtn := d.GetElementByID("lsSaveButton").(*dom.HTMLButtonElement)
	saveBtn.AddEventListener("click", false, func(_ dom.Event) {
		save()
	})

	clearAllBtn := d.GetElementByID("lsClearAllButton").(*dom.HTMLButtonElement)
	clearAllBtn.AddEventListener("click", false, func(_ dom.Event) {
		clear()
	})

	display()
}

func main() {
	d := dom.GetWindow().Document().(dom.HTMLDocument)

	// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
	switch rs := d.ReadyState(); rs {
	case "loading": // HTML parsing
		d.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
			go run()
		})
	// interactive: HTML parsed
	// complete: CSS, image loaded
	case "interactive", "complete":
		run()
	default:
		fmt.Printf("encountered unexpected document ready state %s\n", rs)
	}
}
