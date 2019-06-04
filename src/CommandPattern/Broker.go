package CommandPattern

type Broker struct {
	orders []Order
}

func (b *Broker) TakeOrder(order Order) {
	b.orders = append(b.orders, order)
}

func (b *Broker) PlaceOrders() {
	for _, order := range b.orders {
		order.Execute()
	}
}
