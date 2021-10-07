package main

import (
	"fmt"
	"inventory/service"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("You have to specify an input file")
		return
	}

	inputFileName := args[1] // input.json
	templateFile := "template.xlsx"
	fmt.Println(templateFile)

	borrow, err := service.ParseJsonFile(inputFileName)
	if err != nil {
		log.Fatalf("failed reading input file | %s", err.Error())
	}

	borrow, err = service.IsBorrowValid(borrow)
	if err != nil {
		log.Fatalf("input data is invalid | %s", err.Error())
	}

	fileName := service.GenerateFileName(borrow)
	fmt.Println(fileName)

	xlsx, err := service.GenerateExcelTemplate(fmt.Sprintf("%s.xlsx", fileName))
	if err != nil {
		log.Fatalf("failed generating template | %s\n", err.Error())
	}

	if err := service.FillExcelTemplate(xlsx, borrow); err != nil {
		log.Fatalf("failed filling borrow details | %s", err.Error())
	}

	if err := xlsx.SaveAs(fmt.Sprintf("%s.xlsx", fileName)); err != nil {
		log.Fatalf("failed creating xlsx template | %s", err.Error())
	}
}
