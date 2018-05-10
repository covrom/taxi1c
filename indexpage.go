package main

import (
	"html/template"
	"io"
)

var indexPage *template.Template

func init() {
	indexPage = template.New("index.html")
	indexPage.Parse(`<!DOCTYPE html>
<html lang="ru">

<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>Управляемая форма (демо)</title>
	<link rel="stylesheet" type="text/css" href="assets/skeleton.css">
</head>

<body>
	<div class="container">
		{{.}}
	</div>
</body>

</html>`)
}

func IndexPage(w io.Writer, p interface{}) error {
	return indexPage.Execute(w, p)
}
