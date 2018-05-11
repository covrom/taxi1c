package cforms

import (
	"fmt"
	"html/template"
)

type S string

func (t S) Render() template.HTML {
	return template.HTML(template.HTMLEscapeString(string(t)))
}

type SH string

func (t SH) Render() template.HTML {
	return template.HTML(string(t))
}

type CI struct {
	V interface{}
}

func (t CI) Render() template.HTML {
	return template.HTML(template.HTMLEscapeString(fmt.Sprint(t.V)))
}

func I(v interface{}) *CI {
	rv := &CI{V: v}
	return rv
}

type CP struct {
	CStdForm
}

func (t CP) Render() template.HTML {
	t.Tag = "p"
	return t.RenderTag()
}

func P(items ...CForm) *CP {
	rv := &CP{}
	rv.Content = CItems(items).Render()
	return rv
}

type CDiv struct {
	CStdForm
}

func (t CDiv) Render() template.HTML {
	t.Tag = "div"
	return t.RenderTag()
}

func Div(class, style string, items ...CForm) *CDiv {
	rv := &CDiv{}
	rv.Class = class
	rv.Style = style
	rv.Content = CItems(items).Render()
	return rv
}

type CSpan struct {
	CStdForm
}

func (t CSpan) Render() template.HTML {
	t.Tag = "span"
	return t.RenderTag()
}

func Span(class, style string, items ...CForm) *CSpan {
	rv := &CSpan{}
	rv.Class = class
	rv.Style = style
	rv.Content = CItems(items).Render()
	return rv
}
