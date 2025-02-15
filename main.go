package main

import (
	"fmt"
	"time"
)

type FinanceCategory struct {
	Name  string
	Value float32
}

type HistoryItem struct {
	TypeOfOperation string
	CategoryName    string
	AmountOfMoney   float32
	DateOfOperation time.Time
}

func changeBalance(usrBalance *float32, amountOfMoney float32, categoriesList []FinanceCategory, numberOfCategory int, typeOfOperation string) {
	switch typeOfOperation {
	case "income":
		*usrBalance += amountOfMoney
		fmt.Printf("Add %f to your balance from %s\n", amountOfMoney, categoriesList[numberOfCategory-1].Name)
	case "expanse":
		*usrBalance -= amountOfMoney
		fmt.Printf("Spent %f from your balance to %s\n", amountOfMoney, categoriesList[numberOfCategory-1].Name)
	default:
		fmt.Println("Unknown type of operation")
	}
}

func showAllCategories(categoriesSlice []FinanceCategory) {
	for i := 0; i < len(categoriesSlice); i++ {
		fmt.Printf("%d - %s\n", i+1, categoriesSlice[i].Name)
	}
}

func addToHistory(historySlice *[]HistoryItem, operationType string, categoryOfOperation string, value float32, operationDate time.Time) {
	*historySlice = append(*historySlice, HistoryItem{operationType, categoryOfOperation, value, operationDate})
}

func showHistory(historySlice []HistoryItem) {
	for i := 0; i < len(historySlice); i++ {
		currentItem := historySlice[i]
		fmt.Println("--------------------------------------------")
		fmt.Printf("\nType: %s -- Category name: %s -- Amount of money: %f -- Date: %s\n", currentItem.TypeOfOperation, currentItem.CategoryName, currentItem.AmountOfMoney, currentItem.DateOfOperation.Format("2006-01-02 15:04:05"))
		fmt.Println("--------------------------------------------")
	}
}

func main() {
	var currentBalance float32
	expanseCategories := []FinanceCategory{}
	incomeCategories := []FinanceCategory{}
	historyOfOperations := []HistoryItem{}

	for {
		fmt.Println("")
		fmt.Println("Current Balance: ", currentBalance)
		fmt.Println("")
		fmt.Println("1 - Add income")
		fmt.Println("2 - Add expanse")
		fmt.Println("3 - Add new category")
		fmt.Println("4 - Show History")

		var userChoise int
		fmt.Println("Choose your option: ")
		fmt.Scan(&userChoise)

		switch userChoise {
		case 1:
			var incomeValue float32
			fmt.Println("Add value: ")
			_, incomeErr := fmt.Scan(&incomeValue)
			if incomeErr != nil || incomeValue < 0 {
				fmt.Println("Invalid input. Please enter a positive number.")
				continue
			}

			showAllCategories(incomeCategories)

			var categoryNumber int
			fmt.Println("Choose number of your category: ")
			_, categoryErr := fmt.Scan(&categoryNumber)
			if categoryErr != nil {
				fmt.Println("This category does not exist")
				continue
			}

			changeBalance(&currentBalance, incomeValue, incomeCategories, categoryNumber, "income")

			currentTime := time.Now()

			categoryName := incomeCategories[categoryNumber-1].Name

			addToHistory(&historyOfOperations, "income", categoryName, incomeValue, currentTime)
		case 2:
			var expanseValue float32
			fmt.Println("Add value: ")
			_, expanseErr := fmt.Scan(&expanseValue)
			if expanseErr != nil || expanseValue < 0 {
				fmt.Println("Invalid input. Please enter a positive number.")
				continue
			}

			showAllCategories(expanseCategories)

			var categoryNumber int
			fmt.Println("Choose number of your category: ")
			_, categoryErr := fmt.Scan(&categoryNumber)
			if categoryErr != nil {
				fmt.Println("This category does not exist")
				continue
			}

			changeBalance(&currentBalance, expanseValue, expanseCategories, categoryNumber, "expanse")

			currentTime := time.Now()

			categoryName := expanseCategories[categoryNumber-1].Name

			addToHistory(&historyOfOperations, "expanse", categoryName, expanseValue, currentTime)
		case 3:
			var typeOfCategory int
			fmt.Println("1 - income")
			fmt.Println("2 - expanse")
			fmt.Println("Enter number of your category type: ")
			_, typeOfCategoryErr := fmt.Scan(&typeOfCategory)
			if typeOfCategoryErr != nil {
				fmt.Println("Invalid number of category type. Try again")
				continue
			}

			var nameOfCategory string
			fmt.Println("Enter category name: ")
			_, categoryNameErr := fmt.Scan(&nameOfCategory)
			if categoryNameErr != nil {
				fmt.Println("Invalid name of category")
				continue
			}

			switch typeOfCategory {
			case 1:
				incomeCategories = append(incomeCategories, FinanceCategory{nameOfCategory, 0})
			case 2:
				expanseCategories = append(expanseCategories, FinanceCategory{nameOfCategory, 0})
			}
		case 4:
			showHistory(historyOfOperations)
		default:
			fmt.Printf("\n\nIncorrect type of option. Please choose the correct one.\n\n")
		}
	}
}
