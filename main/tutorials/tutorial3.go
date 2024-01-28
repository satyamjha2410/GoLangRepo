/*
This is tutorial 3 based on standard library used in GO
*/
package main

import (
	"fmt"
	"strings"
	"sort"
)

func main(){

	//Strings has lot of usable functions.
	var name = "Hi there Mr Anthony Gonsalsvis"

	fmt.Println(strings.Contains(name, "Hi"))


	//sort package

	var ages = []int{3,1,9,5,3,90,31}

	sort.Ints(ages)

	// Here original slices are altered
	fmt.Println(ages)


}