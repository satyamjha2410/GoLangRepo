/* Tutorial-1 */
// This tutorial deals with basic Hello world and variable type in go language
/*benifits of golang-
1. Garbage collection is better than c++.
2. It has good concurrency handling.
3. Many features are simplistic to use.
4. GOroutine and Channel are de

*/
package main

import "fmt"

//you can declare a variable outside any function as well

var variable string = "Jojo"

func main(){
	fmt.Println("Hello world")
	
	var name string  = "cent"
	var nameno = 2323

	// used for first time import
	name3 := "afafsd"

	fmt.Println(name)
	fmt.Println(nameno)
	fmt.Println(name, nameno, name3, variable)
}
