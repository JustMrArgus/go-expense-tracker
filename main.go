package main

import "fmt"

type ExpanseCategory struct {
	name  string
	value int
}

type IncomeCategory struct {
	name  string
	value int
}

func addExpanseCategory(usrSlice []ExpanseCategory, categoryName string) {
	newCategory := ExpanseCategory{categoryName, 0}
	usrSlice = append(usrSlice, newCategory)
}

func addIncomeCategory(usrSlice []IncomeCategory, categoryName string) {
	newCategory := IncomeCategory{categoryName, 0}
	usrSlice = append(usrSlice, newCategory)
}

func main() {
	currentBalance := 0.0
	expanseCategories := []ExpanseCategory
	incomeCategoris := []IncomeCategory

	userChoise := 0

	for {
		fmt.Println(currentBalance)
		fmt.Println("")
		fmt.Println("1 - Add income")
		fmt.Println("2 - Add expanse")

	}
}
