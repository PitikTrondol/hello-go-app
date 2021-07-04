package main

import(
	"fmt"
	say_hello "github.com/PitikTrondol/hello-go-module"
)

func main(){
	fmt.Println(say_hello.SayHello())
	fmt.Println(say_hello.SayGoodMorning())
}