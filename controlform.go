package main

import (
	"html/template"

	"github.com/covrom/taxi1c/cforms"
)

func ControlForm() template.HTML {
	cf := cforms.CItems{
		cforms.PHtml(`Сетка является
			<u>12-колоночной подвижной сеткой с максимальной шириной 960px</u>, которая адаптируется под браузер. 
			Максимальная ширина может быть изменена одной строкой CSS и все колонки изменятся одномоментно. 
			Синтаксис прост и позволяет легко создавать адаптивный дизайн страниц. 
			Чтобы увидеть это, измените размер браузера.`),
		&cforms.Div{
			Class: "grayback",
			Style: "text-align:center;",
			Items: cforms.CItems{
				&cforms.Row{Cols: []cforms.Col{
					cforms.Col{
						Width: cforms.W1,
						Items: cforms.CItems{
							cforms.S("Одна"),
						},
					},
					cforms.Col{
						Width: cforms.W11,
						Items: cforms.CItems{
							cforms.S("Одиннадцать"),
						},
					},
				}},
			},
		},
	}
	return cf.Render()
}
