package main

import "fmt"

type FinanceCategory struct {
	name  string
	value float32
}

func changeBalance(usrBalance *float32, amountOfMoney float32, categoriesList []FinanceCategory, numberOfCategory int, typeOfOperation string) {
	switch typeOfOperation {
	case "income":
		*usrBalance += amountOfMoney
	case "expanse":
		*usrBalance -= amountOfMoney
	default:
		fmt.Println("Unknown type of operation")
	}
	categoriesList[numberOfCategory-1].value += amountOfMoney
	fmt.Printf("Add %f to your balance from %s\n", amountOfMoney, categoriesList[numberOfCategory-1].name)
}

func showAllCategories(categoriesSlice []FinanceCategory) {
	for i := 0; i < len(categoriesSlice); i++ {
		fmt.Printf("%d - %s\n", i+1, categoriesSlice[i].name)
	}
}

func main() {
	var currentBalance float32
	expanseCategories := []FinanceCategory{}
	incomeCategories := []FinanceCategory{}

	var userChoise int

	for {
		fmt.Println("")
		fmt.Println(currentBalance)
		fmt.Println("")
		fmt.Println("1 - Add income")
		fmt.Println("2 - Add expanse")
		fmt.Println("3 - Add new category")

		fmt.Println("Choose your option: ")
		fmt.Scan(&userChoise)

		switch userChoise {
		case 1:
			var incomeValue float32
			fmt.Println("Add value: ")
			fmt.Scan(&incomeValue)

			showAllCategories(incomeCategories)

			var categoryNumber int
			fmt.Println("Choose number of your category: ")
			fmt.Scan(&categoryNumber)

			changeBalance(&currentBalance, incomeValue, incomeCategories, categoryNumber, "income")
		case 2:
			var expanseValue float32
			fmt.Println("Add value: ")
			fmt.Scan(&expanseValue)

			showAllCategories(expanseCategories)

			var categoryNumber int
			fmt.Println("Choose number of your category: ")
			fmt.Scan(&categoryNumber)

			changeBalance(&currentBalance, expanseValue, expanseCategories, categoryNumber, "expanse")
		case 3:
			var typeOfCategory int
			fmt.Println("1 - income")
			fmt.Println("2 - expanse")
			fmt.Println("Enter number of your category type: ")
			fmt.Scan(&typeOfCategory)

			var nameOfCategory string
			fmt.Println("Enter category name: ")
			fmt.Scan(&nameOfCategory)

			switch typeOfCategory {
			case 1:
				incomeCategories = append(incomeCategories, FinanceCategory{nameOfCategory, 0})
			case 2:
				expanseCategories = append(expanseCategories, FinanceCategory{nameOfCategory, 0})
			}
		}
	}
}
