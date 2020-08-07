package CommandPattern

type Order interface {
	Execute()
}

type BuyStock struct {
	abcStock *Stock
}

func (b *BuyStock) BuyStock(abcStock *Stock) {
	b.abcStock = abcStock
}

func (b *BuyStock) Execute() {
	b.abcStock.Buy()
}

type SellStock struct {
	abcStock *Stock
}

func (s *SellStock) SellStock(abcStock *Stock) {
	s.abcStock = abcStock
}

func (s *SellStock) Execute() {
	s.abcStock.Sell()
}
