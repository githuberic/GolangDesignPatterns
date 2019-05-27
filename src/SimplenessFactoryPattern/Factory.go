package SimplenessFactoryPattern

import "fmt"

type Factory interface {
	GetShape(shapeType string) (Shape, bool)
}

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

type ShapeFactory struct {
}

func (s *ShapeFactory) GetShape(shapeType string) (Shape, bool) {
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
