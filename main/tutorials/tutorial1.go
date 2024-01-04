/* Tutorial-1 */

// This tutorial deals with basic Hello world and variable type in go language
/*benifits of golang-
1. Garbage collection is better than c++.
2. It has good concurrency handling.
3. Many features are simplistic to use.
4. GOroutine and Channel are decent use for async operations
5. Rich set of standard library for almost all usecases
6. Decent way of dependency management using Go-Mod

*/

// Dependency injection in Golang : https://medium.com/avenue-tech/dependency-injection-in-go-35293ef7b6


/*

FMT package



*/


package main

import "fmt"

//you can declare a variable outside any function as well

var variable string = "Jojo"

func main(){
	fmt.Println("Hello world")
	
	var name string  = "cent"
	var nameno = 2323
	floatvar := 234.2323
	var cameo float64 = 34.3232323

	// used for first time import
	name3 := "afafsd"

	fmt.Println(name)
	fmt.Println(nameno)
	fmt.Println(name, nameno, name3, variable, floatvar, cameo)
	fmt.Printf("My var is %v and second var is %v \n",name,nameno)
	//type of the variable
	fmt.Printf("cameo is of type %T \n", cameo)
	// Precision printing of the type float
	fmt.Printf("here is %0.1f \n", cameo)

	//sprintf

	var formattedString = fmt.Sprintf("My %v is not in the movie",cameo)

	fmt.Println(formattedString)
}
