package main

import "fmt"
import "os"
import "rsc.io/quote"

import "log"

func main(){
	var a int = 10
	var b int32 = 20
	fmt.Println(a+int(b))



	// // explicit
	// var foo string = "foo"

	// // type inferred
	// var bar = "foo"

	// // shorthand
	// baz := "bar"

	// // constant
	// const qux = "qux"


	//float
	var c float32 = 10.34
	fmt.Println(int(c))


	//Boolean
	ok:=true
	fmt.Println(ok)


	m:=1
	if m==1{
		fmt.Println("is true")
	}

	//string
	s1:="hello"
	s2:="tips go"
	s:= `row1\\r\n
	rows2`

	fmt.Println(s1,s2,s)


	//concat

	s3:= s1+s2
	fmt.Println(s3)


	//length
	fmt.Println(len(s3))

	fmt.Println(s3[2:4])


	fmt.Println(quote.Go())

	fmt.Fprintln(os.Stderr, "print to stderr")

	// Package log writes to standard error ánd prints the date and time of each logged message)
	log.Println("hello world")



	///Template String
	name := "bob"
	age := 21
	message := fmt.Sprintf("%s is %d years old", name, age)

	fmt.Println(message)


	// // primitives
	// var myBool bool = true
	// var myInt int = 10
	// var myInt8 int8 = 10
	// var myInt16 int16 = 10
	// var myInt32 int32 = 10
	// var myInt64 int64 = 10
	// var myUint uint = 10
	// var myUint8 uint8 = 10
	// var myUint16 uint16 = 10
	// var myUint32 uint32 = 10
	// var myUint64 uint64 = 10
	// var myUintptr uintptr = 10
	// var myFloat32 float32 = 10.5
	// var myFloat64 float64 = 10.5
	// var myComplex64 complex64 = -1 + 10i
	// var myComplex128 complex128 = -1 + 10i
	// var myString string = "foo"
	// var myByte byte = 10  // alias to uint8
	// var myRune rune = 'a' // alias to int32

	// // composite types
	// var myStruct struct{} = struct{}{}
	// var myArray []string = []string{}
	// var myMap map[string]int = map[string]int{}
	// var myFunction func() = func() {}
	// var myChannel chan bool = make(chan bool)
	// var myInterface interface{} = nil
	// var myPointer *int = new(int)
}