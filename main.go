package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content,err := getNoteData()
	if err != nil{
		fmt.Println(err)
		return
	}
}

func getNoteData()(string,string, error){
	title,err := getUserInput("Note title")
	if err != nil{
		return "","",err
	}
	content ,err := getUserInput("Note content")
	if err != nil{
		return "","",err
	}
	return title, content, err
}

func getUserInput(propt string) (string,error){
	fmt.Println(propt)
	var value string
	fmt.Scanln(&value)
	if value == ""{
		return "", errors.New("invalid empty string")
	}
	return value,nil
}