package main

import "fmt"

func main() {
	fmt.Println("Structs in GO")
	ritik := User{"Ritik", "ritik.go.dev", true, 25}
	fmt.Println(ritik)
	fmt.Printf("details: %+v \n", ritik)
	fmt.Print()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
