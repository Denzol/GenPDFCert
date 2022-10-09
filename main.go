package main

import (
	pdfGen "Golang-HTML-TO-PDF-Converter/pdfGenerator"
	"fmt"
	"log"
)

func main() {
	//Структура для заполнения данных в шаблоне html
	type data struct {
		Title   string
		Student string
		Course  string
		Mentors string
		Date    string
	}

	r := pdfGen.NewRequestPdf("")

	//Выбор шаблона html
	template := "templates/sample.html"

	//Путь размещения сгенерированного файла pdf
	outputPath := "storage/Sertificate.pdf"

	//Данные для заполнения html шаблона
	templateData := data{
		Title:   "Certificate Golang School",
		Student: "Khramtsov Denis",
		Course:  "Become a gopher",
		Mentors: "Pavel Gordiyanov, Mikita Viarbovikau, Sergey Shtripling",
		Date:    "08.09.2022",
	}

	//В переменную "r" записываются данные заполненного шаблона
	err := r.ParseTemplate(template, templateData)
	if err != nil {
		log.Fatal(err)
	}
	//Генерация PDF
	ok, err := r.GeneratePDF(outputPath, "A7")
	if ok {
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
