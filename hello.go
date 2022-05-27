package main

import(
	"fmt"
	. "github.com/yuelinxing/gateio/greetings"

)

func main(){
     message := Hello("hihi")
     fmt.Println(message)
}