package main

import "fmt"


func main() {
	// helloworld
	fmt.Printf("hello,world\n")

	// definition
	my := "My"
	name := "name"
	is := "is"
	Taro := "Taro"
	fmt.Println(my,name,is,Taro)

	//Printf format
	fmt.Printf("17(10s)=%d 17(2s)=%b 17(8s)=%o 17(16s)=%x\n",17,17,17,17)

	//nameless function
	f := func(x,y int)int{return x * y}
	fmt.Println(f(8,9))

	fmt.Printf("%#v\n", func(x,y int)int{return x + y})
	fmt.Printf("%#v\n", func(x,y int)int{return x + y}(14,16))
}
