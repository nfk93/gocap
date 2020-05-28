package capchan

//import "fmt"

type type_func_lp_int_rp_ struct {
	rs      int
	channel (chan func(int) )
	users   []interface{}
}

type Type_func_lp_int_rp_ interface {
	Receive(interface{}) func(int) 
	Send(func(int) , interface{})
	Join(interface{}, interface{})
}

func (c *type_func_lp_int_rp_) Receive(ref interface{}) func(int)  {
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

func (c *type_func_lp_int_rp_) Send(i func(int) , ref interface{}) {
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
func (c *type_func_lp_int_rp_) Join(newuser interface{}, olduser interface{}) {
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


func New_func_lp_int_rp_(rs int, users []interface{}) Type_func_lp_int_rp_ {
	return &type_func_lp_int_rp_{rs, make(chan func(int) ), users}
}