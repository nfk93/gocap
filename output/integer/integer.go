package integer

import fmt "fmt"
import time "time"

type Integer struct { 
i int
flag int
}

func NewInteger(i int,flag int) *Integer {
 return & Integer {i , flag } 
 }

func (v *Integer) Sendcaretaker(c Type__st_Integer) {
 v . Enable ( ) 
 c.Send(v, v) 
 time . Sleep ( 500 * time . Microsecond ) 
 v . Disable ( ) 
 fmt . Println ( "Revoked" ) 
 }

func (v *Integer) Receivecaretaker(c Type__st_Integer) {
 B := c.Receive(v) 
 fmt . Println ( "value of B.i before writing:" , B . read ( ) ) 
 B . write ( 42 ) 
 fmt . Println ( "value of B.i after writing:" , B . read ( ) ) 
 time . Sleep ( 1000 * time . Microsecond ) 
 B . write ( 5 ) 
 fmt . Println ( "value of B.i after revocation:" , B . read ( ) ) 
 }

func (v *Integer) read() int {
 return v . i 
 }

func (v *Integer) write(x int) {
 if v . flag > 0 {
 v . i = x 
 } 
 }

func (v *Integer) Enable() {
 v . flag = 1 
 }

func (v *Integer) Disable() {
 v . flag = 0 
 }


//import "fmt"

type type__st_Integer struct {
	rs      int
	channel (chan *Integer)
	users   []interface{}
}

type Type__st_Integer interface {
	Receive(interface{}) *Integer
	Send(*Integer, interface{})
	Join(interface{}, interface{})
}

func (c *type__st_Integer) Receive(ref interface{}) *Integer {
	valid := false
	//fmt.Printf("[recv] ref= %p \n", ref)
	for _, user := range c.users {
		if user == ref {
			valid = true
		}
	}
	if c.rs <= 1 && valid { //receive from a send only capchan
		ret, _ := <-c.channel
		return ret
	} else {
		panic("Cannot receive: not a user of the channel")
	}
}

func (c *type__st_Integer) Send(i *Integer, ref interface{}) {
	valid := false
	//fmt.Printf("[send] ref= %p \n", ref)
	for _, user := range c.users {
		if user == ref {
			valid = true
		}
	}
	if c.rs >= 1 && valid {
		c.channel <- i
	} else {
		panic("Cannot send: not a user of the channel")
	}
}

//join
func (c *type__st_Integer) Join(newuser interface{}, olduser interface{}) {
	flag := false
	for _, user := range c.users {
		if user == olduser {
			c.users = append(c.users, newuser)
			//fmt.Printf("[join] newuser= %p \n", newuser)
			flag = true
			break
		}
	}
	if !flag {
		panic("Cannot join: not a user of the channel")
	}
}


func New__st_Integer(rs int, users []interface{}) Type__st_Integer {
	return &type__st_Integer{rs, make(chan *Integer), users}
}