package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)


func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}


func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}


func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}


func main(){
	fmt.Println("Hello World")


	/////For
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	///For Short Hand
	sum1 := 1
	for ; sum1 < 1000; {
		sum1 += sum1
	}
	fmt.Println(sum1)

	///For is Go's "while"
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	//Forever
	// for {
	// }	

	////If
	fmt.Println(sqrt(2), sqrt(-4))


	///If with a short statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	////If Else
	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)


	//Switch case
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}


	//Switch evaluation order
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}


	//Switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	///The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	///LIKE AS STACK LIFO
	defer fmt.Println("World1")	
	defer fmt.Println("world") 

	fmt.Println("hello")
}