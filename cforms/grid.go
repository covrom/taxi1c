package cforms

import (
	"html/template"
)

type ColType string

const (
	W1  ColType = "one column"
	W2  ColType = "two columns"
	W3  ColType = "three columns"
	W4  ColType = "four columns"
	W5  ColType = "five columns"
	W6  ColType = "six columns"
	W7  ColType = "seven columns"
	W8  ColType = "eight columns"
	W9  ColType = "nine columns"
	W10 ColType = "ten columns"
	W11 ColType = "eleven columns"
	W12 ColType = "twelve columns"
)

type Grid struct {
	Rows []*Row
}

func (g *Grid) Render() (rv template.HTML) {

	rv = `<div>`
	for _, r := range g.Rows {
		rv += r.Render()
	}
	rv += `</div>`
	return
}

type Row struct {
	Cols []*Col
}

func (r *Row) Render() (rv template.HTML) {

	rv = `<div class="row">`
	for _, c := range r.Cols {
		rv += c.Render()
	}
	rv += `</div>`
	return
}

func ARow(cols ...*Col) *Row {
	return &Row{Cols: cols}
}

type Col struct {
	Width ColType
	Items CItems
}

func (c *Col) Render() (rv template.HTML) {
	rv = `<div class="` + template.HTML(c.Width) + `">`
	if len(c.Width) == 0 {
		rv = `<div style="width:100%;">`
	}
	for _, item := range c.Items {
		rv += item.Render()
	}
	rv += `</div>`
	return
}

func ACol(w ColType, items CItems) *Col {
	return &Col{Width: w, Items: items}
}
