package main

import "fmt"

func main()  {
	fmt.Println("Here We go")
	anything("Chuks",20)

}
func anything(name string, age int) (string, int) {
	age= age + 5
	fmt.Println("my name is ",name)
	fmt.Println("I am", age)

	return name, age
}