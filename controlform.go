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
			f.ARow(
				f.ACol(mwid[i], f.CItems{mcol[i]}),
				f.ACol(mwid[10-i], f.CItems{mcol[10-i]}),
			))
	}

	cf := f.Items(
		f.PHtml(`Сетка является
			<u>12-колоночной подвижной сеткой с максимальной шириной 960px</u>, которая адаптируется под браузер. 
			Максимальная ширина может быть изменена одной строкой CSS и все колонки изменятся одномоментно. 
			Синтаксис прост и позволяет легко создавать адаптивный дизайн страниц. 
			Чтобы увидеть это, измените размер браузера.`),
		f.ADiv("grayback", "text-align:center;", grsmp),
	)
	return cf.Render()
}
