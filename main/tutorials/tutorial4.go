package main

/*
tutorial 4 

For Loop and if-else
*/
import(
	"fmt"
)

func main(){
	
	
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(i)
	}

	names := []string{"mario", "satyam", "Ram" , "Krishna"}

	for i := 0 ;i < len(names) ; i++ {

		fmt.Println(names[i])
	}

	for index , value := range names {
		fmt.Printf("Index : %v Value : %v\n", index , value)
	}

	// IF you dont need the index use _
	for _, value := range names {
		fmt.Printf("Value : %v\n" , value)

		//if I alter the value here it wont affect the array as the value is the local variable for the array
		value = "New String"
	}
	// Still prints the same value
	fmt.Println(names)



	//if else conditional

	for index, value := range names {

		// be very careful why adding if and else as Go add semicolon while compiling . strictly forces to use this format no extra line between if and else
		if (index < 2) {
			fmt.Printf(" Less than 2 %v\n", value)
		} else {
			fmt.Printf("More than 2 %v\n", value)
		}	
	}



}