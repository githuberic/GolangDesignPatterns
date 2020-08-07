package CommandPattern

import "fmt"

type Stock struct {
	Name     string
	Quantity int
}

func NewStock() *Stock {
	stock := new(Stock)
	stock.Name = "ABC"
	stock.Quantity = 10
	return stock
}

func (s *Stock) Buy() {
	fmt.Println("Stock [ Name: ", s.Name, ", Quantity: ", s.Quantity, " ] bought")
}

func (s *Stock) Sell() {
	fmt.Println("Stock [ Name: ", s.Name, ",Quantity: ", s.Quantity, " ] sold")
}
