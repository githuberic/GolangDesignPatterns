package FlyweightPattern

import "fmt"

//创建一个工厂，生成基于给定消息的实体类的对象
type ShapeFactory struct {
	circleMap map[string]*Shape
}

func (s *ShapeFactory) GetCircle(color string) Shape {
	if s.circleMap == nil {
		s.circleMap = make(map[string]*Shape)
	}

	circle := s.circleMap[color]
	if circle == nil {
		circle = &Circle{}
		circle.Circle(color)
		s.circleMap[color] = circle
		fmt.Println("Creating circle of color : ", color)
	}
	return circle
}
