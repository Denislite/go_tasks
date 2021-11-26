package figures

import (
	"math"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
	GREEN
	PURPLE
	GREY
)

type Color byte

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW", "GREEN", "PURPLE", "GREY"}
	return strings[c]
}

type Field struct {
	size   int
	paral  Parallelogram
	circle Circle
	dot    [3]Dot
}

type Dot struct {
	x     float64
	y     float64
	color Color
}

type Parallelogram struct {
	peak [4]Dot
}

type Circle struct {
	peak  Dot
	color Color
	area  float64
}

func (d *Dot) SetColor(color Color) {
	d.color = color
}

func (c *Circle) SetColor(f Field) {
	c.color = Color(byte(c.area / float64(f.size)))
}

func (f *Field) DrawFigure(p Parallelogram) {
	f.paral = p
}

func (f *Field) DrawDots(d [3]Dot) {
	f.dot = d
}

func (f *Field) DrawCircle(c Circle) {
	c.peak = f.paral.FindCenter()
	c.area = f.paral.CalculateArea()
	f.circle = c
}

func (p *Parallelogram) CalculatePeak() Dot {
	center := p.FindCenter()
	x := 2*center.x - p.peak[1].x
	y := 2*center.y - p.peak[1].y
	color := Color((x + y) / 2)
	peak := Dot{x, y, color}
	p.peak[3] = peak
	return peak
}

func (p Parallelogram) CalculateArea() float64 {
	side_a := math.Sqrt(math.Pow((p.peak[0].x-p.peak[1].x), 2) + math.Pow((p.peak[0].y-p.peak[1].y), 2))
	side_b := math.Sqrt(math.Pow((p.peak[1].x-p.peak[2].x), 2) + math.Pow((p.peak[1].y-p.peak[2].y), 2))
	diag := math.Sqrt(math.Pow((p.peak[0].x-p.peak[2].x), 2) + math.Pow((p.peak[0].y-p.peak[2].y), 2))
	halfper := (side_a + side_b + diag) / 2
	area := 2 * math.Sqrt(halfper*(halfper-side_a)*(halfper-side_b)*(halfper-diag))
	return area
}

func (p Parallelogram) FindCenter() Dot {
	x := (p.peak[0].x + p.peak[2].x) / 2
	y := (p.peak[0].y + p.peak[2].y) / 2
	color := Color((x + y) / 2)
	center := Dot{x, y, color}
	return center
}

func (c Circle) FindRadius() float64 {
	radius := math.Sqrt(c.area / math.Pi)
	return math.Round(radius*100) / 100
}

func (c Circle) FindLength() float64 {
	length := c.FindRadius() * math.Pi * 2
	return math.Round(length*100) / 100
}
