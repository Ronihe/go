package shape

type Shape interface {
	Area() float64
	Perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) Area() float64 {
	return r.width * r.height
}

func (r *rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func NewRect(width, height float64) Shape {
	return &rect{width: width, height: height}
}
