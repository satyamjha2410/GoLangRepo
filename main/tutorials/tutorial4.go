package main

/*
tutorial 4 

For Loop
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

	for _, value := range names {
		fmt.Printf("Value : %v\n" , value)

		//if I alter the value here it wont affect the array as the value is the local variable for the array
		value = "New String"
	}
	fmt.Println(names)



}