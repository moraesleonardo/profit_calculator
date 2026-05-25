package main

import (
	"fmt"

)

func main(){

//var revenue float64
//var expenses float64
//var tax_rate float64

//fmt.Print("Revenue: ")
//fmt.Scan(&revenue)
//fmt.Print("Expenses: ")
//fmt.Scan(&expenses)
//fmt.Print("Tax Rate: ")
//fmt.Scan(&tax_rate)

revenue := getUserInput("Revenue: ") 
expenses := getUserInput("Expenses: ") 
tax_rate := getUserInput("Tax Rate:: ") 

earning_before_tax, earnings_after_tax, ratio := calculateFinancials(revenue, expenses, tax_rate)
fmt.Println(earning_before_tax)
fmt.Println(earnings_after_tax)
fmt.Printf("Ratio is: %.2f\n", ratio)
}

func calculateFinancials(revenue, expenses, tax_rate float64)(float64, float64, float64){
	earning_before_tax := revenue - expenses
	earnings_after_tax := earning_before_tax * (1 - tax_rate / 100)
	ratio := earning_before_tax / earnings_after_tax
	return earning_before_tax, earnings_after_tax, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}