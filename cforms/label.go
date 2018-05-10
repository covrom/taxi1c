package cforms

import "html/template"

type S string

func (t S) Render() (rv template.HTML) {
	return template.HTML(template.HTMLEscapeString(string(t)))
}

type P string

func (t P) Render() (rv template.HTML) {
	return `<p>` + template.HTML(template.HTMLEscapeString(string(t))) + `</p>`
}

type PHtml string

func (t PHtml) Render() (rv template.HTML) {
	return `<p>` + template.HTML(string(t)) + `</p>`
}

type Div struct {
	Class string
	Style string
	Items CItems
}

func (d Div) Render() (rv template.HTML) {
	rv = `<div`
	if len(d.Class) > 0 {
		rv += ` class="` + template.HTML(d.Class) + `"`
	}
	if len(d.Style) > 0 {
		rv += ` style="` + template.HTML(d.Style) + `"`
	}
	rv += `>`
	for _, item := range d.Items {
		rv += item.Render()
	}
	rv += `</div>`
	return
}

type Span struct {
	Class string
	Style string
	Items CItems
}

func (d Span) Render() (rv template.HTML) {
	rv = `<span`
	if len(d.Class) > 0 {
		rv += ` class="` + template.HTML(d.Class) + `"`
	}
	if len(d.Style) > 0 {
		rv += ` style="` + template.HTML(d.Style) + `"`
	}
	rv += `>`
	for _, item := range d.Items {
		rv += item.Render()
	}
	rv += `</span>`
	return
}
