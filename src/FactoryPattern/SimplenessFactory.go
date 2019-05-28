package FactoryPattern

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (s *Rectangle) Draw() {
	fmt.Print("draw Rectangle!")
}

type Square struct {
}

func (s *Square) Draw() {
	fmt.Println("draw Square!")
}

type SimplenessFactory struct {
}

func (s *SimplenessFactory) GetShape(shapeType string) (Shape, bool) {
	if shapeType == "" {
		return nil, false
	}

	switch shapeType {
	case "Rectangle":
		return new(Rectangle), true
	case "Square":
		return new(Square), true
	default:
		return nil, false
	}
}
