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

func (t *CP) SetId(id string) *CP {
	t.Id = id
	return t
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

func (t *CDiv) SetId(id string) *CDiv {
	t.Id = id
	return t
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

func (t *CSpan) SetId(id string) *CSpan {
	t.Id = id
	return t
}

func Span(class, style string, items ...CForm) *CSpan {
	rv := &CSpan{}
	rv.Class = class
	rv.Style = style
	rv.Content = CItems(items).Render()
	return rv
}

type CH1 struct {
	CStdForm
}

func (t CH1) Render() template.HTML {
	t.Tag = "h1"
	return t.RenderTag()
}

func (t *CH1) SetId(id string) *CH1 {
	t.Id = id
	return t
}

func H1(items ...CForm) *CH1 {
	rv := &CH1{}
	rv.Content = CItems(items).Render()
	return rv
}

type CH2 struct {
	CStdForm
}

func (t CH2) Render() template.HTML {
	t.Tag = "h2"
	return t.RenderTag()
}

func (t *CH2) SetId(id string) *CH2 {
	t.Id = id
	return t
}

func H2(items ...CForm) *CH2 {
	rv := &CH2{}
	rv.Content = CItems(items).Render()
	return rv
}

type CH3 struct {
	CStdForm
}

func (t CH3) Render() template.HTML {
	t.Tag = "h3"
	return t.RenderTag()
}

func (t *CH3) SetId(id string) *CH3 {
	t.Id = id
	return t
}

func H3(items ...CForm) *CH3 {
	rv := &CH3{}
	rv.Content = CItems(items).Render()
	return rv
}

type CH4 struct {
	CStdForm
}

func (t CH4) Render() template.HTML {
	t.Tag = "h4"
	return t.RenderTag()
}

func (t *CH4) SetId(id string) *CH4 {
	t.Id = id
	return t
}

func H4(items ...CForm) *CH4 {
	rv := &CH4{}
	rv.Content = CItems(items).Render()
	return rv
}

type CH5 struct {
	CStdForm
}

func (t CH5) Render() template.HTML {
	t.Tag = "h5"
	return t.RenderTag()
}

func (t *CH5) SetId(id string) *CH5 {
	t.Id = id
	return t
}

func H5(items ...CForm) *CH5 {
	rv := &CH5{}
	rv.Content = CItems(items).Render()
	return rv
}

type CH6 struct {
	CStdForm
}

func (t CH6) Render() template.HTML {
	t.Tag = "h6"
	return t.RenderTag()
}

func (t *CH6) SetId(id string) *CH6 {
	t.Id = id
	return t
}

func H6(items ...CForm) *CH6 {
	rv := &CH6{}
	rv.Content = CItems(items).Render()
	return rv
}

type CUpCase struct {
	CStdForm
}

func (t CUpCase) Render() template.HTML {
	t.Tag = "span"
	cls := "upperspaced"
	if len(t.Class) > 0 {
		t.Class += " " + cls
	} else {
		t.Class = cls
	}
	return t.RenderTag()
}

func UpCase(items ...CForm) *CUpCase {
	rv := &CUpCase{}
	rv.Content = CItems(items).Render()
	return rv
}

type CCode struct {
	CStdForm
}

func (t CCode) Render() template.HTML {
	t.Tag = "code"
	return t.RenderTag()
}

func (t *CCode) SetId(id string) *CCode {
	t.Id = id
	return t
}

func Code(text string, items ...CForm) *CCode {
	rv := &CCode{}
	rv.Content = template.HTML(template.HTMLEscapeString(text)) + CItems(items).Render()
	return rv
}
