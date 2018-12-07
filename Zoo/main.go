package main

// animals
import(
	"fmt"
	"./animals"


)

func main(){
	fmt.Println(Appname())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
	fmt.Println(animals.LionFeed())
	fmt.Println(animals.DolphinFeed())

}