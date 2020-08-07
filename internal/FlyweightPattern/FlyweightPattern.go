package FlyweightPattern

import "fmt"

type Shape interface {
	Draw3()
	SetX(x int)
	SetY(y int)
	SetRadius(radius int)
}

type Circle struct {
	color  string
	x      int
	y      int
	radius int
}

func (c *Circle) Circle(color string) {
	c.color = color
}

func (c *Circle) Draw3() {
	fmt.Println("Circle: Draw() [Color :", c.color, "x:", c.x, "y:", c.y, "radius:", c.radius)
}
func (c *Circle) SetX(x int) {
	c.x = x
}
func (c *Circle) SetY(y int) {
	c.y = y
}
func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}
