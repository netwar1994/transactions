package card

import "fmt"

func TranslateMCC(code string) string {
	mcc := map[string]string{
		"5411": "Grocery Stores, Supermarkets",
		"5812": "Eating Places and Restaurants",
	}
	for key, _ := range mcc {
		if key == code {
			fmt.Println(mcc[key])
			return mcc["key"]
		}
	}
	return "Category not listed"
}