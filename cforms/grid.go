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

type CGrid struct {
	Rows []*CRow
	CStdForm
}

func (g *CGrid) Render() template.HTML {
	g.Tag = "div"
	g.Content = ""
	for _, r := range g.Rows {
		g.Content += r.Render()
	}
	return g.RenderTag()
}

func (t *CGrid) SetId(id string) *CGrid {
	t.Id = id
	return t
}

func Grid(rows ...*CRow) *CGrid {
	return &CGrid{Rows: rows}
}

type CRow struct {
	Cols []*CCol
	CStdForm
}

func (r *CRow) Render() template.HTML {
	r.Tag = "div"
	r.Content = ""
	cls := "row"
	if len(r.Class) > 0 {
		r.Class += " " + cls
	} else {
		r.Class = cls
	}
	for _, c := range r.Cols {
		r.Content += c.Render()
	}
	return r.RenderTag()
}

func (t *CRow) SetId(id string) *CRow {
	t.Id = id
	return t
}

func Row(cols ...*CCol) *CRow {
	return &CRow{Cols: cols}
}

type CCol struct {
	Width ColType
	CStdForm
}

func (c *CCol) Render() template.HTML {

	c.Tag = "div"
	cls := string(c.Width)
	if len(c.Width) == 0 {
		cls = string(W12)
	}
	if len(c.Class) > 0 {
		c.Class += " " + cls
	} else {
		c.Class = cls
	}
	return c.RenderTag()
}

func (t *CCol) SetId(id string) *CCol {
	t.Id = id
	return t
}

func Col(w ColType, items ...CForm) *CCol {
	rv := &CCol{Width: w}
	rv.Content = CItems(items).Render()
	return rv
}
