package NullObjectPattern

type AbstractCustomer interface {
	IsNil() bool
	GetName() string
}

type RealCustomer struct {
	name string
}

func (r *RealCustomer) RealCustomer(name string) AbstractCustomer {
	r.name = name
	return r
}

func (r *RealCustomer) IsNil() bool {
	return false
}

func (r *RealCustomer) GetName() string {
	return r.name
}

type NullCustomer struct {
	name string
}

func (n *NullCustomer) IsNil() bool {
	return true
}

func (n *NullCustomer) GetName() string {
	return "Not Available in Customer Database"
}
