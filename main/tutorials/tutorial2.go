/*
Tutorial : 2
This start for arrays and slices

*/
package main

import (
	"fmt"
)
	

func main(){
	fmt.Println("Hello")

	// Three ways to initialize an array
	var ages[3] int = [3] int {12,3,4}

	var ages2 = [3] int{1,2,3}

	ages3 := [3] int {2,3,4}

	fmt.Println(ages, ages2, ages3)


	//For slices we dont have to mention the size
	var score = []int{1,2,3}

	score = append(score, 89)
	//append operation returns a new array

	scoreSubset := score[1:4]

	score2 := []int{}

	score2 = append(score2, 12)

	fmt.Println(score, scoreSubset)

	fmt.Println(score2, scoreSubset)

	
	

}


