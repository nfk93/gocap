package main

import fmt "fmt"
import "github.com/nfk93/gocap/output/capchan"

type test struct { 
i int
}

func (ts *test) foo(ch capchan.Type_int) {
 tmp := ch.Receive(ts) 
 ts . i = tmp 
 fmt . Println ( ts . i ) 
 }

func main() {
 ch := capchan.New_int(1, [](interface{}){capchan.TopLevel}) 
 v1 := 42 
 ts := test {0 } 
 ts_ := & ts 
 ch.Join(ts_, capchan.TopLevel) 
 go ts . foo ( ch ) 
 ch.Send(v1, capchan.TopLevel) 
 }

