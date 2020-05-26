package main

import (
	fmt "fmt"

	"github.com/nfk93/gocap/output/capchan"
)

func main() {
	ch := capchan.New_int(1, [](interface{}){capchan.TopLevel})
	v1 := 123
	v2 := ch.Receive(capchan.TopLevel)
	ch.Send(v1, capchan.TopLevel)
	fmt.Println(v2)
}
