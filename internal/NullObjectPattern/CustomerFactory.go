package NullObjectPattern

type CustomerFactory struct {
	Names []string
}

func (c *CustomerFactory) GetCustomer(nameP string) AbstractCustomer {
	for _, name := range c.Names {
		if name == nameP {
			return new(RealCustomer).RealCustomer(name)
		}
	}
	return new(NullCustomer)
}
