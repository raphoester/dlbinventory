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
		fmt.Println("help")
		return
	}

	filename := args[1] // input.json

	borrow, err := service.ParseJsonFile(filename)
	if err != nil {
		log.Fatalf("failed reading input file | %s", err.Error())
	}

	borrow, err = service.IsBorrowValid(borrow)
	if err != nil {
		log.Fatalf("input data is invalid | %s", err.Error())
	}

	fileName := service.GeneratePdfFileName(borrow)
	fmt.Println(fileName)
}
