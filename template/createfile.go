package template

import (
	"os"
	"text/template"
)

func Iar_eww(name string, path string, data IarEwwType) {
	tmpl, err := template.ParseFiles("template/raw/iar/iar-eww.txt")
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create(path + name + ".eww")
	if err != nil {
		panic(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(outputFile)

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		panic(err)
	}
}

func Iar_ewp(name string, path string, data IarEwpType) {
	tmpl, err := template.ParseFiles("template/raw/cc2530/cc2530-iar-ewp.txt")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(path+"/"+name, os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(path+"/"+name+"/code", os.ModePerm)
	if err != nil {
		panic(err)
	}

	Iar_main(name, path)

	outputFile, err := os.Create(path + "/" + name + "/" + name + ".ewp")
	if err != nil {
		panic(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(outputFile)

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		panic(err)
	}
}

func Iar_main(name string, path string) {
	tmpl, err := template.ParseFiles("template/raw/cc2530/cc2530-iar-main.c.txt")
	if err != nil {
		panic(err)
	}
	outputFile, err := os.Create(path + "/" + name + "/code/main.c")
	if err != nil {
		panic(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(outputFile)

	err = tmpl.Execute(outputFile, nil)
	if err != nil {
		panic(err)
	}
}
