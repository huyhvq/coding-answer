package triangle

type Type int

const (
	Scalene     Type = iota + 1
	Isosceles
	Equilateral
	Error
)

func GetTriangleType(a, b, c int) (Type) {
	if isTheSidesOfATriangle(a, b, c) {
		if a == b && a == c {
			return Equilateral
		} else if a == b || a == c || b == c {
			return Isosceles
		}
		return Scalene
	}
	return Error
}

func isTheSidesOfATriangle(a, b, c int) (bool) {
	if a > 0 && b > 0 && c > 0 && a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}
