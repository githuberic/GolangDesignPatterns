package SimplenessFactoryPattern

import "fmt"

type Factory interface {
	GetShape(shapeType string) (Shape, bool)
}

type Shape interface {
	draw()
}

type Rectangle struct {
}

func (s Rectangle) draw() {
	fmt.Print("draw Rectangle!")
}

type Square struct {
}

func (s Square) draw() {
	fmt.Println("draw Square!")
}

type ShapeFactory struct {
}

func (s *ShapeFactory) GetShape(shapeType string) (Shape, bool) {
	if shapeType == "" {
		return nil, false
	}

	switch shapeType {
	case "Rectangle":
		return Rectangle{}, true
	case "Square":
		return Square{}, true
	default:
		return nil, false
	}
}
