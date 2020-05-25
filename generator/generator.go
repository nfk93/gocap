package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nfk93/gocap/parser/simple/ast"
	"github.com/nfk93/gocap/parser/simple/token"
)

type Attrib interface{}

var typesArray []string

// var typesArray = []string{"string", "hello_type"}

var makeNewCapChannelTemplate = "capchan.New_$TYPE(1, [](interface{}){$USER})"
var sendCapChannelTemplate = "$CHAN.Send($VAL, $USER)"
var receiveCapChannelTemplate = "$CHAN.Receive($USER)"
var joinCapChannelTemplate = "$CHAN.Join($NUSER, $USER)"
var typeCapChannelTemplate = "capchan.Type_$TYPE"

func remove_Parentheses(typeString string) string {
	//typeString = strings.ReplaceAll(typeString, " ", "_s_")
	typeString = strings.ReplaceAll(typeString, "(", "_lp_")
	typeString = strings.ReplaceAll(typeString, ")", "_rp_")
	return typeString
}

//CapChanMake: Typ: string, VarId :string
//return: string
func MakeNewCapChannelType(c *ast.CapChanMake) (interface{}, error) {
	typeString := c.Typ
	typeStringU := remove_Parentheses(typeString)
	receiverString := c.VarId

	result := strings.Replace(makeNewCapChannelTemplate, "$TYPE", typeStringU, -1)
	result = strings.Replace(result, "$USER", receiverString, -1)

	fmt.Printf("[make capchan]type: %s, receiver: %s\n", typeString, receiverString)
	fmt.Println("[make capchan]generated code: ", result)

	flag := false
	for _, t := range typesArray {
		if t == typeStringU {
			flag = true
			break
		}
	}
	if !flag {
		typesArray = append(typesArray, typeStringU)
	}

	return result, nil
}

//path to executable(used for compile)
func getPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	return exPath
}

//path to codefile(used for run/test)
func getPath2() string {
	_, filename, _, _ := runtime.Caller(1)
	exPath := filepath.Dir(filename)
	fmt.Println(exPath)
	return exPath
}

func createFile(data string, filename string) {
	//TODO
	currentPathString := getPath2()
	packageDirString := currentPathString + "/capchan"
	if _, err := os.Stat(packageDirString); os.IsNotExist(err) {
		err_ := os.Mkdir(packageDirString, 0777)
		if err_ != nil {
			panic("Cannot create dirctory ./capchan")
		}
	}
	filePathString := packageDirString + "/" + filename
	f, err := os.OpenFile(filePathString, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
		//panic("Cannot open file " + filePathString)
	}
	_, err = f.Write([]byte(data))
	if err != nil {
		panic("Cannot write file " + filePathString)
	}
	f.Close()
}

func NewCapChannelPackage() {
	data, err := ioutil.ReadFile(getPath2() + "/template")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	tempString := string(data)
	for _, typeStringU := range typesArray {
		dataString := strings.ReplaceAll(tempString, "$TYPE", typeStringU)
		filenameString := "capchan_" + typeStringU + ".go"
		createFile(dataString, filenameString)
	}
}

func SendCapChannel(receiverString string, c *ast.CapChanSend) (interface{}, error) {
	channelString := c.ChannelId
	valueString := c.SendId

	result := strings.Replace(makeNewCapChannelTemplate, "$CHAN", channelString, -1)
	result = strings.Replace(result, "$VAL", valueString, -1)
	result = strings.Replace(result, "$USER", receiverString, -1)

	return result, nil
}

func ReceiveCapChannel(c *ast.CapChanReceive) (interface{}, error) {
	channelString := c.ChannelId
	receiverString := c.ReceiverId

	result := strings.Replace(makeNewCapChannelTemplate, "$CHAN", channelString, -1)
	result = strings.Replace(result, "$USER", receiverString, -1)

	return result, nil
}

//TODO: no struct for join now
func JoinCapChannel(channel Attrib, newuser Attrib, receiver Attrib) (interface{}, error) {
	channelString := string(channel.(*token.Token).Lit)
	newuserString := string(newuser.(*token.Token).Lit)
	receiverString := string(receiver.(*token.Token).Lit)

	result := strings.Replace(makeNewCapChannelTemplate, "$CHAN", channelString, -1)
	result = strings.Replace(result, "$NUSER", newuserString, -1)
	result = strings.Replace(result, "$USER", receiverString, -1)

	return result, nil
}

//TODO: only accept a string without whitespace
func TypeCapChannel(typeString string) (interface{}, error) {
	typeString = remove_Parentheses(typeString)

	result := strings.Replace(typeCapChannelTemplate, "$TYPE", typeString, -1)

	return result, nil
}

func MakeNewChannelType(typ Attrib) (interface{}, error) {
	fmt.Println("found chan type: ", string(typ.(*token.Token).Lit))
	return nil, nil
}
