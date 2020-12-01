package darts

//Points received based on circle.
const (
	outsideTarget = 0
	outerCircle   = 1
	middleCircle  = 5
	innerCircle   = 10
)

//Score returns earned points in a single toss of a darts game.
func Score(x, y float64) (result int) {
	//The result is based on where the dart lands within each concentric circle.
	switch {
	//If the dart lands on the border it is within inner circle.
	case (x*x)+(y*y) <= 1*1:
		result = innerCircle
	case (x*x)+(y*y) <= 5*5 && (x*x)+(y*y) > 1*1:
		result = middleCircle
	case (x*x)+(y*y) <= 10*10 && (x*x)+(y*y) > 5*5:
		result = outerCircle
	case (x*x)+(y*y) > 10*10:
		result = outsideTarget
	}
	return result
}
