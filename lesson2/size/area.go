package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi*math.Pow(c.Radius, 2)
}

type Square struct {
	Width float64
	Height float64
}

func (s *Square) Area() float64 {
	return s.Width * s.Height
}

type Sizer interface {
	Area() float64
}

type Shape interface {
	Sizer
	fmt.Stringer
}

func main() {
	c := Circle{Radius: 10}
	fmt.Println(c)

	s := Square{Width: 10, Height: 8}
	fmt.Println(s)

	// Eq(c,s)

	//for loop types
	i := 0
	for i <= 3 {
		i += 1
	}

	for j := 7; j < 10; j++ {
		
	}

	// for { //infinite loop

	// }

	switch i {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	}


}

func Eq(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	} else {
		return s2
	}
}



