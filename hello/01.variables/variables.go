package main

import "fmt"


///Global variables
var d int

var Pi float32 = 3.14

// func getMoney() (int,int){
// 	return 100,1500
// }

func main(){
	var a = 40
	var b = 50

	//Khai bao bien c
	c :=a+b
	fmt.Println("c::",c)


	//Global variables,local variables 
	//If global and local variables has the same name ,it grab the local variables
	var Pi int  = 4
	fmt.Printf("a=%d, b=%d, c=%d, d=%d\n",a,b,c,d)
	fmt.Println("Pi:::",Pi)
}

// func main(){
// 	///var name type
// 	var userName ="Tu";
// 	println("UserName:",userName)
// 	fmt.Printf("Type of user: ̀%T\n", userName)


// 	///Short name,can not reasinged value when using short name
// 	///The scope variable is used in function can not declare out of funcition
// 	age :=40
// 	fmt.Println("Age:",age)


// 	// //error:can not reasinged value when using short name
// 	// //Khai bao bien email
// 	// var email = "tu1@gmail.com"
// 	// //khai bao lai use short name
// 	// email :="tu123@gmail.com"


// 	///Asign multiple value
// 	var a = 1
// 	var b = 2

// 	b,a = a,b
// 	fmt.Println(a,b)
	

// 	///_ can not use
// 	m1, _:=getMoney()
// 	_,m2 :=getMoney()

// 	fmt.Println("m1::",m1)
// 	fmt.Println("m2::",m2)
// }