package cforms

import "html/template"

type CForm interface {
	Render() template.HTML
}

type CItems []CForm

func (ci CItems) Render() (rv template.HTML) {
	for _, i := range ci {
		rv += i.Render()
	}
	return
}

func Items(items ...CForm) CItems {
	return items
}

type CStdForm struct {
	Tag     string
	Id      string
	Class   string
	Style   string
	Attrs   string
	Content template.HTML
}

func (sf *CStdForm) RenderTag() (res template.HTML) {
	res = "<" + template.HTML(sf.Tag)
	if len(sf.Id) > 0 {
		res += ` id="` + template.HTML(sf.Id) + `"`
	}
	if len(sf.Class) > 0 {
		res += ` class="` + template.HTML(sf.Class) + `"`
	}
	if len(sf.Style) > 0 {
		res += ` style="` + template.HTML(sf.Style) + `"`
	}
	if len(sf.Attrs) > 0 {
		res += " " + template.HTML(sf.Attrs)
	}
	res += ">" + sf.Content + "</" + template.HTML(sf.Tag) + ">"
	return
}

