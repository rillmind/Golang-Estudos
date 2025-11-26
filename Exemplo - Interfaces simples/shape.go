package main

import "fmt"

type shape interface {
	area() float32
}

type triangle struct {
	base   float32
	height float32
}

type square struct {
	base   float32
	height float32
}

func printArea(s shape) {
	fmt.Printf("A área do formato é %v\n", s.area())
}

func (t triangle) area() float32 {
	return (t.base * t.height) / 2
}

func (s square) area() float32 {
	return s.base * s.height
}
