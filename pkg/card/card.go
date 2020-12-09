package card

import (
	"fmt"
)

type Transaction struct {
	Id int64
	Sum int64
	Date string
	MCC string
	Status string
}

type Card struct {
	Id int64
	Issuer string
	Balance int64
	Currency string
	Number string
	Transactions []Transaction
}

func Info(card *Card) {
	fmt.Println("\n------------ Information about card -------------")
	fmt.Println("Issuer: ",card.Issuer)
	fmt.Println("Balance:", card.Balance)
	fmt.Println("Card number:" ,card.Number)
	fmt.Println("Currency:", card.Currency)
	fmt.Println("\n----------------- Transactions ------------------")
	for _, transaction := range card.Transactions {
		fmt.Println("Transaction ID:", transaction.Id)
		fmt.Println("Transaction sum:", transaction.Sum)
		fmt.Println("Transaction date:", transaction.Date)
		fmt.Println("MMC:", transaction.MCC)
		fmt.Println("Transaction status:", transaction.Status)
		fmt.Println("------------------------------------------------")
	}
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Balance -= transaction.Sum
	fmt.Println("\n------------------------------------------------")
	fmt.Println("Transaction date:", transaction.Date)
	fmt.Println("Merchant category code:", transaction.MCC)
	fmt.Println("Transaction sum:", transaction.Sum)
	fmt.Println("Transaction status:", transaction.Status)
}

func SumByMMC(transactions []Transaction, mcc []string) int64 {
	sum := int64(0)
	for _, transaction := range transactions {
		for _, m := range mcc {
			if transaction.MCC == m {
				sum += transaction.Sum
			}
		}
	}
	return sum
}
