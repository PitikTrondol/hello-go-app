package main

import(
	"fmt"
	say_hello "github.com/PitikTrondol/hello-go-module/v2"
)

func main(){
	fmt.Println(say_hello.SayHello("Tekek"))
	fmt.Println(say_hello.SayGoodMorning())
}