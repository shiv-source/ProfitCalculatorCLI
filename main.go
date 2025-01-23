package main

import "fmt"

func main() {
	var (
		revenue  float64
		expenses float64
		taxRate  float64
	)

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate (in %): ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	fmt.Printf("Earnings Before Tax (EBT): %.2f\n", earningBeforeTax)

	if earningBeforeTax > 0 {
		earningAfterTax := earningBeforeTax - ((earningBeforeTax * taxRate) / 100)
		fmt.Printf("Profit After Tax: %.2f\n", earningAfterTax)

		if expenses > 0 {
			profitPercentage := (earningAfterTax / expenses) * 100
			fmt.Printf("Profit Percentage: %.2f%%\n", profitPercentage)
		} else {
			fmt.Println("Expenses are zero or negative, profit percentage cannot be calculated.")
		}
	} else {
		fmt.Println("No profit to calculate after tax, as earnings before tax are zero or negative.")
	}
}
