package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

func main() {
	funcionarios := ReadCsv()
	for _, funcionario := range funcionarios {
		println(funcionario.Departamento, funcionario.Nome)
	}
}

type Funcionario struct {
	Departamento string `csv:"Departamento"`
	Nome         string `csv:"Funcionário"`
}

func ReadCsv() []*Funcionario {
	csvFile, csvFileError := os.OpenFile("funcionarios.csv", os.O_RDWR, os.ModePerm)
	if csvFileError != nil {
		panic(csvFileError)
	}
	defer csvFile.Close()

	var funcionarios []*Funcionario
	if unmarshalError := gocsv.UnmarshalFile(csvFile, &funcionarios); unmarshalError != nil {
		panic(unmarshalError)
	}

	return funcionarios
}
