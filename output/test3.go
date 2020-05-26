package main

import fmt "fmt"
import "github.com/nfk93/gocap/output/capchan"

func foo(ch capchan.Type_int) {
 v2 := ch.Receive() 
 fmt . Println ( v2 ) 
 }

func main() {
 ch := capchan.New_int(1, [](interface{}){}) 
 v1 := 42 
 go foo ( ch ) 
 ch.Send(v1, ) 
 }

h.Send(v1, capchan.TopLevel)
}
