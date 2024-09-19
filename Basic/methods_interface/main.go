package main

import (
	"fmt"
	"math"
)

// Methods
type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

// 	Remember: a method is just a function with a receiver argument.

// Here's Abs written as a regular function with no change in functionality.


// Pointer receivers
v.Scale(10)
	fmt.Println(v.Abs())

}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

	
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)

	//Methods and pointer indirection
	// var v4 Vertex
	// ScaleFunc(v4, 5)  // Compile error!
	// ScaleFunc(&v4, 5) // OK

	// var v Vertex
	// v.Scale(5)  // OK
	// p := &v
	// p.Scale(10) // OK

	// v := Vertex{3, 4}
	// v.Scale(2)
	// ScaleFunc(&v, 10)

	// p := &Vertex{4, 3}
	// p.Scale(3)
	// ScaleFunc(p, 8)

	// fmt.Println(v, p)
	// {60 80} &{96 72}

	//Methods and pointer indirection (2)

	// func (v Vertex) Abs() float64 {
	// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
	// }
	
	// func AbsFunc(v Vertex) float64 {
	// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
	// }
	// var v Vertex
	// fmt.Println(AbsFunc(v))  // OK
	// fmt.Println(AbsFunc(&v)) // Compile error!

	// var v Vertex
	// fmt.Println(v.Abs()) // OK
	// p := &v
	// fmt.Println(p.Abs()) // OK



	//////////////Interfaces
	// var a Abser
	// f := MyFloat(-math.Sqrt2)
	// v := Vertex{3, 4}

	// a = f  // a MyFloat implements Abser
	// a = &v // a *Vertex implements Abser

	// // In the following line, v is a Vertex (not *Vertex)
	// // and does NOT implement Abser.
	// a = v

	// fmt.Println(a.Abs())


	///The empty interface
	// var i interface{}
	// describe(i)

	// i = 42
	// describe(i)

	// i = "hello"
	// describe(i)

	// fmt.Printf("(%v, %T)\n", i, i)

	// (<nil>, <nil>)
	// (42, int)
	// (hello, string)

	//////Type assertions
	// var i interface{} = "hello"

	// s := i.(string)
	// fmt.Println(s)

	// s, ok := i.(string)
	// fmt.Println(s, ok)

	// f, ok := i.(float64)
	// fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)


	///Stringers ovveride toString
	// func (p Person) String() string {
	// 	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
	// }

}

// func describe(i interface{}) {
// 	fmt.Printf("(%v, %T)\n", i, i)
// }


// type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }


func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

///Type switches check type
// func do(i interface{}) {
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	case string:
// 		fmt.Printf("%q is %v bytes long\n", v, len(v))
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)
// 	}

// }

// do(21)
// do("hello")
// do(true)

