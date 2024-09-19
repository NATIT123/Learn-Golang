package main

import (
	"fmt"
	"math"
	"strings"
)

func printSlice(s []int) {
	//A nil slice has a length and capacity of 0 and has no underlying array.
	if s == nil {
		fmt.Println("nil!")
	}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main(){
	fmt.Println("Hello World")


	////Pointer
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	fmt.Println(*p) // read j through the pointer
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j	


	///Struct
	// Uppercase package : can use another packages
	// lowecase package : can not use
	type Vertex struct {
		X int
		Y int
	}

	fmt.Println(Vertex{1, 2})

	//Struct Fields
	v := Vertex{1, 2}
	v.X = 4
	v.Y=3
	fmt.Println(v.X)


	//Pointers to structs
	po:= &v
	po.X = 1e9
	fmt.Println(v)



	///Struct Literals
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		h  = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, *h, v2, v3)


	//Arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)


	var s []int = primes[1:4]
	fmt.Println(s)

	///Slices are like references to arrays
	///Pointer to array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, c)
	fmt.Println(names)


	///Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	l := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(l)


	//Slice Defaults
	sl := []int{2, 3, 5, 7, 11, 13}

	var s1 = sl[1:4]
	s1[0]=5
	fmt.Println(s1)

	s = sl[:2]
	fmt.Println(s)

	s = sl[1:]
	fmt.Println(sl)

	// 	The length of a slice is the number of elements it contains.

	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.


	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	///Creating a slice with make
	// a := make([]int, 5)  // len(a)=5
	// b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	// b = b[:cap(b)] // len(b)=5, cap(b)=5
	// b = b[1:]      // len(b)=4, cap(b)=4

	z := make([]int, 5)
	printSlice1("z", z)

	x := make([]int, 0, 5)
	printSlice1("x", x)

	y := x[:2]
	printSlice1("y", y)

	m := y[2:5]
	printSlice1("m", m)


	///Slices of slices

	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}


	///Appending to a slice
	var sa []int
	var temp []int
	printSlice(sa)

	// append works on nil slices.
	sa = append(sa, 0)
	printSlice(sa)
	printSlice(temp)

	// The slice grows as needed.
	sa = append(sa, 1)
	printSlice(sa)

	// We can add more than one element at a time.
	sa = append(sa, 2, 3, 4)


	////return a new slice and can not modify slide
	temp = append(sa, 4)
	printSlice(temp)
	printSlice(sa)


	//Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}


	//Range cotinued
	pow1 := make([]int, 10)
	for i := range pow {
		pow1[i] = 1 << uint(i) // == 2**i
	}

	///_ represent for value not used
	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}

	///Maps
	type Vertex1 struct {
		Lat, Long float64
	}
	var m1 map[string]Vertex1= make(map[string]Vertex1)
	m1["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}


	///2222
	// var m = map[string]Vertex{
	// 	"Bell Labs": {40.68433, -74.39967},
	// 	"Google":    {37.42202, -122.08408},
	// }
	
	fmt.Println(m1["Bell Labs"])


	////Mutating Maps
	m2 := make(map[string]int)

	m2["Answer"] = 42
	fmt.Println("The value:", m2["Answer"])

	m2["Answer"] = 48
	fmt.Println("The value:", m2["Answer"])

	delete(m2, "Answer")
	fmt.Println("The value:", m2["Answer"])


	///Function values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))


	

}

func printSlice1(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}	

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

