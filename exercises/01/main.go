package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)
func getArgs()([]string,error)  {
	args:= os.Args[1:]
	if len(args)<3{
		return nil,errors.New("invalid args")
	}
	result:= []string{
		args[0],
		args[1],
		strings.Join(args[2 : len(args)-1], " "),
		args[len(args)-1],
	}

	return result,nil
}

func getName(firstname string,lastname string,middlename string, country_code string)( string, bool){
	switch country_code{

	case "US":
		return strings.Join([]string{firstname,middlename,lastname}, " "),false
	case "VN":
		 return strings.Join([]string{lastname,middlename,firstname}, " "),false
	default:
		 return strings.Join([]string{lastname,middlename,firstname}, " "),true
	}
}
func main()   {
    args, error := getArgs()
	if error!= nil{
		fmt.Println(error)
		return
	}
    name, isDefaultFormat := getName(args[0],args[1],args[2],args[3])
	fmt.Printf("Output: %s ",name)

	if isDefaultFormat{
		fmt.Printf("(default format: VN)\n")
	}else{	
		fmt.Printf("(format: %s)\n",args[3])
	}

}