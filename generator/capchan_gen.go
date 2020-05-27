package generator

import (
	"fmt"
	"os"
	"strings"

	"github.com/nfk93/gocap/utils"
)

type Attrib interface{}

var typesArray []string

// var typesArray = []string{"string", "hello_type"}

var makeNewCapChannelTemplate = "capchan.New_$TYPE(1, [](interface{}){$USER})"
var sendCapChannelTemplate = "$CHAN.Send($VAL, $USER)"
var receiveCapChannelTemplate = "$CHAN.Receive($USER)"
var joinCapChannelTemplate = "$CHAN.Join($NUSER, $USER)"
var packageCapChannelTemplate = `package capchan

//import "fmt"

type type_$TYPEU struct {
	rs      int
	channel (chan $TYPE)
	users   []interface{}
}

type Type_$TYPEU interface {
	Receive(interface{}) $TYPE
	Send($TYPE, interface{})
	Join(interface{}, interface{})
}

func (c *type_$TYPEU) Receive(ref interface{}) $TYPE {
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

func (c *type_$TYPEU) Send(i $TYPE, ref interface{}) {
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
func (c *type_$TYPEU) Join(newuser interface{}, olduser interface{}) {
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


func New_$TYPEU(rs int, users []interface{}) Type_$TYPEU {
	return &type_$TYPEU{rs, make(chan $TYPE), users}
}`

//CapChanMake: Typ: string, VarId :string
//return: string
func MakeNewCapChannelType(typeString, receiverString string) string {
	typeStringU := utils.RemoveParentheses(typeString)

	result := strings.Replace(makeNewCapChannelTemplate, "$TYPE", typeStringU, -1)
	result = strings.Replace(result, "$USER", receiverString, -1)

	fmt.Printf("[make capchan]type: %s, receiver: %s\n", typeString, receiverString)
	fmt.Println("[make capchan]generated code: ", result)

	flag := false
	for _, t := range typesArray {
		if t == typeString {
			flag = true
			break
		}
	}
	if !flag {
		typesArray = append(typesArray, typeString)
	}

	return result
}

func createPackage(data string, filename string, output string) {
	currentPathString := output
	//TODO
	packageDirString := currentPathString + "/capchan"
	if _, err := os.Stat(packageDirString); os.IsNotExist(err) {
		err_ := os.Mkdir(packageDirString, 0777)
		if err_ != nil {
			panic("Cannot create dirctory ./capchan")
		}
	}
	filePathString := packageDirString + "/" + filename
	CreateFile(data, filePathString)
}

func CreateFile(data string, filepath string) {
	filePathString := filepath
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

func GenerateCapChannelPackage(outputPath string) {
	tempString := packageCapChannelTemplate
	for i, typeString := range typesArray {
		typeStringU := utils.RemoveParentheses(typeString)
		dataString := strings.ReplaceAll(tempString, "$TYPEU", typeStringU)
		dataString = strings.ReplaceAll(dataString, "$TYPE", typeString)
		if i == 0 {
			dataString += "\nconst TopLevel = \"LBS\""
		}
		filenameString := "capchan_" + typeStringU + ".go"
		if utils.IfPrintPackages {
			printPackages(filenameString, dataString)
		} else {
			createPackage(dataString, filenameString, outputPath)
		}
	}
}

func printPackages(filenameString, dataString string) {
	fmt.Printf("[generator]: ===== target file: %s =====\n", filenameString)
	fmt.Printf("[generator]: \n %s \n", dataString)
}

func SendCapChannel(channelString, valueString, receiverString string) string {

	result := strings.Replace(sendCapChannelTemplate, "$CHAN", channelString, -1)
	result = strings.Replace(result, "$VAL", valueString, -1)
	result = strings.Replace(result, "$USER", receiverString, -1)

	return result
}

func ReceiveCapChannel(channelString, receiverString string) string {

	result := strings.Replace(receiveCapChannelTemplate, "$CHAN", channelString, -1)
	result = strings.Replace(result, "$USER", receiverString, -1)

	return result
}

func JoinCapChannel(channelString, newuserString, receiverString string) string {

	result := strings.Replace(joinCapChannelTemplate, "$CHAN", channelString, -1)
	result = strings.Replace(result, "$NUSER", newuserString, -1)
	result = strings.Replace(result, "$USER", receiverString, -1)

	return result
}
