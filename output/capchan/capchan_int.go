package capchan

import "fmt"

type type_int struct {
	rs      int
	channel (chan int)
	users   []interface{}
}

type Type_int interface {
	Receive(interface{}) int
	Send(int, interface{})
	Join(interface{}, interface{})
}

func (c *type_int) Receive(ref interface{}) int {
	valid := false
	fmt.Printf("[recv] ref= %p \n", ref)
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

func (c *type_int) Send(i int, ref interface{}) {
	valid := false
	fmt.Printf("[send] ref= %p \n", ref)
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
func (c *type_int) Join(newuser interface{}, olduser interface{}) {
	flag := false
	for _, user := range c.users {
		if user == olduser {
			c.users = append(c.users, newuser)
			fmt.Printf("[join] newuser= %p \n", newuser)
			flag = true
			break
		}
	}
	if !flag {
		panic("Cannot join: not a user of the channel")
	}
}


func New_int(rs int, users []interface{}) Type_int {
	return &type_int{rs, make(chan int), users}
}

const TopLevel = "LBS"