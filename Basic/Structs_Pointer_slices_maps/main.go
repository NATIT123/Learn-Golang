package main

import "fmt"

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
}