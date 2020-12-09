package main

import (
	"github.com/netwar1994/transactions/pkg/card"
	"fmt"
	"time"
)

func main()  {
	mccList := []string{"5411", "5812"}

	transaction1 := card.Transaction{
		Id:     0,
		Sum:    1000,
		Date:   time.Now().Format(time.Stamp),
		MCC:    "5411",
		Status: "pending",
	}
	transaction2 := card.Transaction{
		Id:     1,
		Sum:    2000,
		Date:   time.Now().Format(time.Stamp),
		MCC:    "5812",
		Status: "pending",
	}
	transaction3 := card.Transaction{
		Id:     2,
		Sum:    22000,
		Date:   time.Now().Format(time.Stamp),
		MCC:    "5813",
		Status: "pending",
	}

	transactions := []card.Transaction{transaction1, transaction2, transaction3}

	master := card.Card{
		Id:           0,
		Issuer:       "MasterCard",
		Balance:      50_000,
		Currency:     "RUB",
		Number:       "1234-5678-1234-5678",
		Transactions: transactions,
	}
	card.AddTransaction(&master, &transaction1)
	card.AddTransaction(&master, &transaction2)
	card.AddTransaction(&master, &transaction3)

	card.Info(&master)
	fmt.Println("Amount of transaction for MCC 5411, 5812:", card.SumByMMC(transactions, mccList))

	category := card.TranslateMCC(master.Transactions[1].MCC)
	fmt.Println(category)
}
