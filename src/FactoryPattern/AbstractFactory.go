package FactoryPattern

import "fmt"

type AbstractFactory interface {
	GetColor(colorType string) Color
	GetShape(shapeType string) Shape
}

type AbsFactory struct {
}

func (s *AbsFactory) GetShape(colorType string) Shape {
	if colorType == "" {
		return nil
	}

	switch colorType {
	case "Rectangle":
		return new(Rectangle)
	case "Square":
		return new(Square)
	default:
		return nil
	}
}

func (s *AbsFactory) GetColor(shapeType string) Color {
	if shapeType == "" {
		return nil
	}

	switch shapeType {
	case "Red":
		return new(Red)
	case "Green":
		return new(Green)
	case "Blue":
		return new(Blue)
	default:
		return nil
	}
}

type Color interface {
	Fill()
}

type Red struct {
}

func (s *Red) Fill() {
	fmt.Println("Red Fill")

}

type Green struct {
}

func (s *Green) Fill() {
	fmt.Println("Green Fill")

}

type Blue struct {
}

func (s *Blue) Fill() {
	fmt.Println("BlueFill")
}
