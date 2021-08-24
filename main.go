package main

import "fmt"

func main()  {
	fmt.Println("Here We go")
	anything("Chuks",20)
	b,t,s := fruits( "Chuks has", 25,"oranges")
	fmt.Println(b,t,s)
	 newAge:=calculateAge(25,10)
	 fmt.Println(newAge)
	 array()

}
func anything(name string, age int) (string, int) {
	age= age + 5
	fmt.Println("my name is ",name)
	fmt.Println("I am", age)

	return name, age
}
func fruits (name string, number int, name2 string)(string,int,string){
	number= number *10
	return  name, number,name2

}
func calculateAge(currentAge,numOfYears int)int{
	result:= currentAge + numOfYears
	return result
}
func array(){
	arr :=[]int{17,26,10,80}
	for idx, val:= range arr{
		fmt.Println(idx,val)
	}

}