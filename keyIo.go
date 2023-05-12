package keyIo

//keyIo returns a string of data from Stdin and nil on success or empty string and nil
//on failure, package user must supply directive to user, i.e. please enter foo etc

import (
	"bufio"
	"os"
)

func GetData()(string,error){
	reader := bufio.NewReader(os.Stdin)
	data,err := reader.ReadString('\n')
	if err != nil{
		return "",err
	}
	return data,nil
}
