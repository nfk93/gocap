package main

import (
	fmt "fmt"
	time "time"

	"github.com/nfk93/gocap/output/capchan"
)

type Integer struct {
	I    int
	Flag int
}

func (v *Integer) Sendcaretaker(c capchan.Type_func_lp__rp_int, d capchan.Type_func_lp_int_rp_) {
	r := v.read
	w := v.write
	v.Enable()
	c.Send(r, v)
	d.Send(w, v)
	time.Sleep(500 * time.Microsecond)
	v.Disable()
	fmt.Print("true value:")
	fmt.Println(v.I)
}

func (v *Integer) Receivecaretaker(c capchan.Type_func_lp__rp_int, d capchan.Type_func_lp_int_rp_) {
	read := c.Receive(v)
	write := d.Receive(v)
	fmt.Println(read())
	write(-1)
	fmt.Println(read())
	time.Sleep(1000 * time.Microsecond)
	write(-2)
	fmt.Println(read())
}

func (v *Integer) read() int {
	if v.Flag > 0 {
		return v.I
	}
	return 0
}

func (v *Integer) write(x int) {
	if v.Flag > 0 {
		v.I = x
	}
}

func (v *Integer) Enable() {
	v.Flag = 1
}

func (v *Integer) Disable() {
	v.Flag = 0
}

type Obj struct {
	B Integer
	C Integer
}

func (A Obj) test() {
	c := capchan.New_func_lp__rp_int(1, [](interface{}){A})
	d := capchan.New_func_lp_int_rp_(1, [](interface{}){A})
	r := &A.C
	c.Join(r, A)
	d.Join(r, A)
	r = &A.B
	c.Join(r, A)
	d.Join(r, A)
	go A.B.Sendcaretaker(c, d)
	go A.C.Receivecaretaker(c, d)
	time.Sleep(3000 * time.Microsecond)
}

func main() {
	A := Obj{Integer{1, 0}, Integer{0, 0}}
	A.test()
}
