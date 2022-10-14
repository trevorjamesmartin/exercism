package clock

import (
	"fmt"
	"math"
)

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	// zero hour
	c := Clock{
		h: 0,
		m: 0,
	}
	// adjust minutes
	switch {
	case m > 0:
		c = c.Add(m)
		break
	case m < 0:
		c = c.Subtract(0 - m)
		break
	}
	// adjust hours
	switch {
	case h > 0:
		c = c.Add(h * 60)
		break
	case h < 0:
		c = c.Subtract((0 - h) * 60)
		break
	}
	return c
}

func (c Clock) Add(m int) Clock {
	if m < 0 {
		return c.Subtract(0 - m)
	}
	var hours = c.h + int(math.Floor(float64(m/60)))
	var minutes = c.m + m%60
	// roll-over ?
	for minutes > 60 {
		hours += 1
		minutes -= 60
	}
	return Clock{h: hours % 24, m: minutes}
}

func (c Clock) Subtract(m int) Clock {
	if m < 0 {
		return c.Add(0 - m)
	}
	var hours = c.h - int(math.Floor(float64(m/60)))
	var minutes = c.m - m%60
	// roll-under ?
	for minutes < 0 {
		hours -= 1
		minutes += 60
	}
	for hours < 0 {
		hours += 24
	}
	return Clock{h: hours, m: minutes}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
