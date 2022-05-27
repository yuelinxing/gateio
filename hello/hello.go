package main

import(
	"fmt"
        "github.com/yuelinxing/gateio/greetings"	
)

func main(){
     message := greetings.Hello("hihi")
     fmt.Println(message)
}