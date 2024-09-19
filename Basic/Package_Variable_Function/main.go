package main

import (
	"fmt"
	"math"
	"math/cmplx"

	"golang.org/x/exp/rand"
)

func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}


func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}



func main(){
	fmt.Println("Hello World")

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi)


	//Functions
	fmt.Println(add(3,2))

	fmt.Println(add1(42, 13))

	///Mutiple return results
	a,b :=swap("Hello","World")
	fmt.Println(a,b)


	//Named return values
	fmt.Println(split(17))


	//Variables
	var i int
	var c, python, java = true, false, "no!"
	k := 3
	fmt.Println(i, c, python, java,k)


	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)


	//Zero values
	// 	0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings.
	var e int
	var f float64
	var g bool
	var s string
	fmt.Printf("%v %v %v %q\n", e, f, g, s)


	// Type conversions
	var x, y int = 3, 4
	var j float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(j)
	fmt.Println(x, j, z)


	////Constant
	const Pi = 3.14
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)


	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	

}