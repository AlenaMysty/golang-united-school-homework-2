package square

import "math"

type number int64

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum number) float64 {
	var s float64 = 0
	if sidesNum == SidesTriangle {
		s = sideLen * sideLen / 4 * math.Sqrt(3)
	} else if sidesNum == SidesSquare {
		s = sideLen * sideLen
	} else if sidesNum == SidesCircle {
		s = math.Pi * sideLen * sideLen
	}

	return s
}
