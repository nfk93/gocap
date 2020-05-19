package generator

import "fmt"
import "lbs_proj/token"

type Attrib interface{}

func DoSomething(s Attrib) (string, error) {
  fmt.Println(string(s.(*token.Token).Lit))
  return string(s.(*token.Token).Lit), nil
}
