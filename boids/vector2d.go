package main

import "math"

// Vector2D A vector for location and velocity
type Vector2D struct {
	x float64
	y float64
}

// Add the vectors
func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x + v2.x, y: v1.y + v2.y}
}

// Subtract the vectors
func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x - v2.x, y: v1.y - v2.y}
}

// Multiply the vectors
func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{x: v1.x * v2.x, y: v1.y * v2.y}
}

// AddV add the value to each vector component
func (v1 Vector2D) AddV(v float64) Vector2D {
	return Vector2D{x: v1.x + v, y: v1.y + v}
}

// SubtractV subtract the value from each vector component
func (v1 Vector2D) SubtractV(v float64) Vector2D {
	return Vector2D{x: v1.x - v, y: v1.y - v}
}

// MultiplyV multiply the vector components by the value
func (v1 Vector2D) MultiplyV(v float64) Vector2D {
	return Vector2D{x: v1.x * v, y: v1.y * v}
}

// DivideV divide the vector components by the value
func (v1 Vector2D) DivideV(v float64) Vector2D {
	return Vector2D{x: v1.x / v, y: v1.y / v}
}

// Limit the vector components to the value ranges
func (v1 Vector2D) Limit(lower, upper float64) Vector2D {
	return Vector2D{x: math.Min(math.Max(v1.x, lower), upper), y: math.Min(math.Max(v1.y, lower), upper)}
}

// Distance gets Euclidean distance bewtween to vectors
func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
