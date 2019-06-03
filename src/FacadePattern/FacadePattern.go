package FacadePattern

import "fmt"

type Shape interface {
	Draw2()
}

type Rectangle struct {
}

func (r *Rectangle) Draw2() {
	fmt.Println("Rectangle::draw()")
}

type Square struct {
}

func (s *Square) Draw2() {
	fmt.Println("Square::draw()")
}

type Circle struct {
}

func (c *Circle) Draw2() {
	fmt.Println("Circle::draw()")
}
