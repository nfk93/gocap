package main

import fmt "fmt"
import time "time"
import "github.com/nfk93/gocap/output/capchan"

type Integer struct { 
I int
}

func (v *Integer) Sendref(c capchan.Type__st_int) {
 fmt . Print ( "Value before send reference to channel:" ) 
 fmt . Println ( v . I ) 
 a := & v . I 
 c.Send(a, v) 
 time . Sleep ( 1000 * time . Microsecond ) 
 fmt . Print ( "One second after send my reference to channel:" ) 
 fmt . Println ( v . I ) 
 }

func (v *Integer) Receiveref(c capchan.Type__st_int) int {
 a := c.Receive(v) 
 * a = v . I 
 return * a 
 }

type Obj struct { 
B Integer
C Integer
}

func (A *Obj) test() {
 c := capchan.New__st_int(1, [](interface{}){A}) 
 ref := & A . B 
 c.Join(ref, A) 
 ref = & A . C 
 c.Join(ref, A) 
 go A . B . Sendref ( c ) 
 go A . C . Receiveref ( c ) 
 time . Sleep ( 3000 * time . Microsecond ) 
 }

func main() {
 A := Obj {Integer {0 } , Integer {1 } } 
 A . test ( ) 
 }

