package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nfk93/gocap/parser/simple/token"
	"github.com/nfk93/gocap/utils"
)

type Attrib interface{}

var typesArray []string

// var typesArray = []string{"string", "hello_type"}

var makeNewCapChannelTemplate = "capchan.New_$TYPE(1, [](interface{}){$USER})"
var sendCapChannelTemplate = "$CHAN.Send($VAL, $USER)"
var receiveCapChannelTemplate = "$CHAN.Receive($USER)"
var joinCapChannelTemplate = "$CHAN.Join($NUSER, $USER)"

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

func GenerateCapChannelPackage(path string) {
	data, err := ioutil.ReadFile(path + "/template")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	tempString := string(data)
	for i, typeString := range typesArray {
		typeStringU := utils.RemoveParentheses(typeString)
		dataString := strings.ReplaceAll(tempString, "$TYPEU", typeStringU)
		dataString = strings.ReplaceAll(dataString, "$TYPE", typeString)
		if i == 0 {
			dataString += "\nconst topLevel = \"LBS\""
		}
		filenameString := "capchan_" + typeStringU + ".go"
		if utils.IfPrintPackages {
			printPackages(filenameString, dataString)
		} else {
			createFile(dataString, filenameString)
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

//TODO: no support to join now
func JoinCapChannel(channel Attrib, newuser Attrib, receiver Attrib) (interface{}, error) {
	channelString := string(channel.(*token.Token).Lit)
	newuserString := string(newuser.(*token.Token).Lit)
	receiverString := string(receiver.(*token.Token).Lit)

	result := strings.Replace(makeNewCapChannelTemplate, "$CHAN", channelString, -1)
	result = strings.Replace(result, "$NUSER", newuserString, -1)
	result = strings.Replace(result, "$USER", receiverString, -1)

	return result, nil
}
