package FlyweightPattern

import (
	"math/rand"
)

type FlyweightPatternDemo struct {
	colors []string
}

func (f *FlyweightPatternDemo) GetRandomColor() string {
	if f.colors == nil {
		f.colors = []string{"Red", "Green", "Blue", "White", "Black"}
	}

	num := rand.Intn(len(f.colors) - 1)
	return f.colors[num]
}

func (f *FlyweightPatternDemo) GetRandomXAndY() int {
	num := rand.Intn(10) * 100
	return num
}
