package main

import "fmt"

// Declare the Expense struct here
// your code goes here

// Implement the Total method to calculate the total amount spent
// your code goes here

// Implement the getName method on the Expense struct here
// your code goes here
type Expense struct {
	name   string
	amount float64
	date   string
}

func Total(expenses []Expense) float64 {
	var total float64 = 0.0
	for _, expense := range expenses {
		total += expense.amount
	}
	return total
}

func (e Expense) getName() string {
	return e.name
}

func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}
