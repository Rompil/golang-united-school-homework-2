package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type SideNum int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesCircle  SideNum = iota
	_
	_
	SidesTriangle 
	SidesSquare 
)

func circleArea(radius float64) (area float64) {
	area = math.Pi * radius * radius
	return
}

func traingleArea(sideLen float64) (area float64) {
	// Herone formular
	semiperimeter := 3.0 * sideLen / 2.0
	area = math.Sqrt(semiperimeter * math.Pow((semiperimeter-sideLen), 3.0))
	return
}

func squareArea(sideLen float64) (area float64) {
	area = sideLen * sideLen
	return
}

func zeroArea(float64) float64 {
	return 0
}

func CalcSquare(sideLen float64, sidesNum SideNum) float64 {
	var f func(float64) float64
	switch int(sidesNum) {
	case 0:
		f = circleArea
	case 3:
		f = traingleArea
	case 4:
		f = squareArea
	default:
		f = zeroArea
	}
	return f(sideLen)
}
