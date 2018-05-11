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
		f.H6(f.UpCase(f.S("Сетка"))).SetId("grid"),
		f.P(f.SH(`Сетка является
			<u>12-колоночной подвижной сеткой с максимальной шириной 960px</u>, которая адаптируется под браузер. 
			Максимальная ширина может быть изменена одной строкой CSS и все колонки изменятся одномоментно. 
			Синтаксис прост и позволяет легко создавать адаптивный дизайн страниц. 
			Чтобы увидеть это, измените размер браузера.`)),
		f.Div("grayback", "text-align:center;", grsmp),
		f.H6(f.UpCase(f.S("Типографика"))).SetId("typography"),
		f.P(f.S(`Все метрики указываются в единицах`),
			f.Code(`rem`), f.S(`, поэтому, размер шрифта и расстояния между элементами гибко адаптируются, основываясь на одной настройке шрифта в`),
			f.Code("<html>"), f.S(`. Из коробки, шаблон никогда не изменяет размер шрифта в`),
			f.Code("<html>"), f.S(`, но Вы можете изменить это для своего проекта. 
			Все размеры используют в качестве базы 10px, например,`),
			f.Code("<h1>"), f.S(`с`), f.Code("5.0rem"), f.S(`размером означает`),
			f.Code("50px"), f.S(".")),
		f.Row(
			f.Col(f.W6,
				f.P(f.SH(` <strong>Основу типографики</strong> составляет шрифт
					<a href="http://www.google.com/fonts/specimen/Noto+Sans">Noto Sans</a> разработанный Google, 
					его размер используется 13rem (13px) при высоте строки в 1.6 (24px). Различные типы текста выглядят так:
					<a href="#">ссылка</a>,
					<strong>строгий</strong>,
					<em>выделенный</em>, и
					<u>подчеркнутый</u>.`)),
				f.P(f.SH(`<strong>Заголовки</strong> обединены в набор, отличающийся параметрами`),
					f.Code("letter-spacing"), f.S(","),
					f.Code("line-height"), f.S(", и"),
					f.Code("margins"), f.S(".")),
				f.Col(f.W6,
					f.H1(f.Code("<h1>"),
						f.S(`50rem`)),
					f.H2(f.Code("<h2>"),
						f.S(`42rem`)),
					f.H3(f.Code("<h3>"),
						f.S(`36rem`)),
					f.H4(f.Code("<h4>"),
						f.S(`30rem`)),
					f.H5(f.Code("<h5>"),
						f.S(`24rem`)),
					f.H6(f.Code("<h6>"),
						f.S(`15rem`))),
			)),
	)
	return cf.Render()
}
