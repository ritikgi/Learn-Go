package main

import "fmt"

func main() {
	fmt.Println("welcome to Slicse")
	var fruitList = []string{"Apple", "Peach", "Banana"}
	fruitList = append(fruitList, "Mango", "Grapes")
	fmt.Println("items in the list %T \n", fruitList)

	fmt.Println(fruitList[:1])
}
