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
