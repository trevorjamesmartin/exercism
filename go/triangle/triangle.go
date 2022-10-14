// Determine if a triangle is equilateral, isosceles, or scalene.
package triangle

// KindFromSides() returns this type.
type Kind string

// joins 3 sides of arbitrary proportions with an identifier
type Triangle struct {
	a          float64
	b          float64
	c          float64
	identifier Kind
}

// determine if this is a triangle
func (t *Triangle) Valid() bool {
	switch {
	case t.a <= 0 || t.b <= 0 || t.c <= 0:
		return false
	case t.a+t.b < t.c:
		return false
	case t.b+t.c < t.a:
		return false
	case t.c+t.a < t.b:
		return false
	}
	return true
}

const (
	NaT = "not a triangle" // not a triangle
	Equ = "equilateral"    // equilateral
	Iso = "isosceles"      // isosceles
	Sca = "scalene"        // scalene
)

// returns the Kind of triangle formed by joining sides a, b, & c
func KindFromSides(a, b, c float64) Kind {
	abc := Triangle{a: a, b: b, c: c}
	switch {
	case !abc.Valid():
		abc.identifier = NaT
	case a != b && b != c && a != c:
		abc.identifier = Sca
	case a == b && b == c && c == a:
		abc.identifier = Equ
	default:
		abc.identifier = Iso
	}
	return abc.identifier
}
