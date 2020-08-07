package DecoratorPattern

import "fmt"

type Shape interface {
	Draw1()
}

type Rectangle struct {
}

func (r *Rectangle) Draw1() {
	fmt.Println("Shape: Rectangle")
}

type Circle struct {
}

func (c *Circle) Draw1() {
	fmt.Println("Shape: Circle")
}

type ShapeDecorator struct {
	decoratedShape Shape
}

func (s *ShapeDecorator) ShapeDecorator(decoratedShape Shape) {
	s.decoratedShape = decoratedShape
}

func (s *ShapeDecorator) Draw1() {
	s.decoratedShape.Draw1()
}

type RedShapeDecorator struct {
	shapeDecorator ShapeDecorator
}

func (s *RedShapeDecorator) RedShapeDecorator(decoratedShape Shape) {
	s.shapeDecorator.ShapeDecorator(decoratedShape)
}

func (s *RedShapeDecorator) Draw1() {
	s.shapeDecorator.Draw1()
	s.setRedBorder(s.shapeDecorator.decoratedShape)
}

func (s *RedShapeDecorator) setRedBorder(decoratedShape Shape) {
	fmt.Println("Border Color: Red")
}
