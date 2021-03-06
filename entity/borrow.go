package entity

import "time"

type Borrow struct {
	Headphones Headphones
	Airbox     Airbox
	Charger    Charger
	Laptop     Laptop
	Mobile     Mobile
	Borrower   Person
	Date       time.Time
	Type       string
}

type Person struct {
	Firstname  string
	Lastname   string
	Department string
}

type Headphones struct {
	ModelName string
	Serial    string
}

type Airbox struct {
	LineNumber string
	ImeiNumber string
}

type Laptop struct {
	SerialNumber string
	ModelName    string
}

type Charger struct {
	ModelName string
}

type Mobile struct {
	LineNumber string
	ModelName  string
	ImeiNumber string
}
