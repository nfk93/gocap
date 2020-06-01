package main

import time "time"
import integer "github.com/nfk93/gocap/output/integer"
import capchan "github.com/nfk93/gocap/output/capchan"

func main() {
 B := integer . NewInteger ( 7 , 0 ) 
 C := integer . NewInteger ( 0 , 0 ) 
 c := integer.New__st_Integer(1, [](interface{}){capchan.TopLevel}) 
 c.Join(B, capchan.TopLevel) 
 c.Join(C, capchan.TopLevel) 
 go B . Sendcaretaker ( c ) 
 go C . Receivecaretaker ( c ) 
 time . Sleep ( 3000 * time . Microsecond ) 
 }

