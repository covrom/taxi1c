package main

import (
	"html/template"

	f "github.com/covrom/taxi1c/cforms"
)

func ControlForm() template.HTML {

	mcol := []f.S{
		"Одна", "Две", "Три", "Четыре", "Пять", "Шесть", "Семь", "Восемь", "Девять", "Десять", "Одиннадцать",
	}
	mwid := []f.ColType{
		f.W1, f.W2, f.W3, f.W4, f.W5, f.W6, f.W7, f.W8, f.W9, f.W10, f.W11, f.W12,
	}

	var grsmp f.CItems
	for i := 0; i < 11; i++ {
		grsmp = append(grsmp,
			f.Row(
				f.Col(mwid[i], mcol[i]),
				f.Col(mwid[10-i], mcol[10-i]),
			))
	}

	cf := f.Items(
		f.P(f.SH(`Сетка является
			<u>12-колоночной подвижной сеткой с максимальной шириной 960px</u>, которая адаптируется под браузер. 
			Максимальная ширина может быть изменена одной строкой CSS и все колонки изменятся одномоментно. 
			Синтаксис прост и позволяет легко создавать адаптивный дизайн страниц. 
			Чтобы увидеть это, измените размер браузера.`)),
		f.Div("grayback", "text-align:center;", grsmp),
	)
	return cf.Render()
}
