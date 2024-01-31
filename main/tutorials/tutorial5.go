package main

/*Tutorial 5
Functions in GOLang

In Go style we can pass function as an argument easily in another method
*/
import(
	"fmt"
	"math"
)

func sayAllowed(st string){
	fmt.Printf("You are welcome %v\n", st)
}

func sayBlocked(st string){
	fmt.Printf("You are blocked %v\n", st)
}

// Passing function as an argument
func repeatNames(narr []string , f func(string)){
	for _, value := range narr{
		f(value)
	}
}

// function returning variable
func areaCircle(r float64 ) float64{
	return math.Pi * r * r;
}

//this function returns 2 values
func lengthAndInitials(names []string) ([]string, int){

	nameInit := []string{}
	// var nameInit []string
	for _, value := range names{
		nameInit = append(nameInit, value[:1])
	}
	return nameInit,len(names)
}

func main(){
	
	names := []string{"Ram", "Krishna", "Shiva", "Hanuman"}
	repeatNames(names, sayAllowed)

	fmt.Println(areaCircle(19.3), areaCircle(10))
	// to print upto 3 places of decimal
	fmt.Printf("Area are %0.3f and %0.3f\n",areaCircle(19.3), areaCircle(10))

	namesInitials, length := lengthAndInitials(names)

	fmt.Println(namesInitials, length)


}