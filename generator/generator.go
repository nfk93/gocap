package generator

import "fmt"
import "github.com/nfk93/gocap/token"

type Attrib interface{}

func DoSomething(s Attrib) (string, error) {
  fmt.Println(string(s.(*token.Token).Lit))
  return string(s.(*token.Token).Lit), nil
}

func MakeNewCapChannelType(typ Attrib) (interface{}, error) {
  fmt.Println("found capchan type: ", string(typ.(*token.Token).Lit))
  return nil, nil
}

func MakeNewChannelType(typ Attrib) (interface{}, error) {
  fmt.Println("found chan type: ", string(typ.(*token.Token).Lit))
  return nil, nil
}
