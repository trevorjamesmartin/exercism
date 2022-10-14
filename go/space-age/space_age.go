package space

type Planet string

type Orbits map[Planet]float64

var EarthYear = Orbits{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	p, ok := EarthYear[planet]
	if !ok {
		return -1
	}
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	years := days / 365.25
	return years / p
}
