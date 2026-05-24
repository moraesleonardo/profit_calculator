package main

import (
	"fmt"

)

func main(){

var revenue float64
var expenses float64
var tax_rate float64

fmt.Print("Revenue: ")
fmt.Scan(&revenue)
fmt.Print("Expenses: ")
fmt.Scan(&expenses)
fmt.Print("Tax Rate: ")
fmt.Scan(&tax_rate)

earning_before_tax := revenue - expenses
earnings_after_tax := earning_before_tax + (earning_before_tax * (tax_rate / 100))
ratio := earning_before_tax / earnings_after_tax

fmt.Println(earning_before_tax)
fmt.Println(earnings_after_tax)
fmt.Printf("Ratio is: %.2f\n", ratio)
}
