package goarea

import "math"

// Pi is the ratio of a circles circumference to its diameter
const Pi = 3.1416

// Circle is responsible for calculating the area of the circumference.
func Circle(radius float64) float64 {
	return math.Pow(radius, 2) * Pi
}

// Rectangle is responsible for calculating the area of a rectangle
func Rectangle(base, height float64) float64 {
	return base * height
}

// not visible
func _TriangleEq(base, height float64) float64 {
	return (base * height) / 2
}
