package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
 userNote,err:=	note.New(title,content)

 if err!= nil{
	fmt.Println(err)
 }
 userNote.Display()
 err= userNote.Save()
 if err != nil{
	fmt.Println(err)
	return
 }
 fmt.Println("file save")
}

func getNoteData()(string,string){
	title := getUserInput("Note title")
	
	content := getUserInput("Note content")
	
	return title, content
}

func getUserInput(propt string) string{
	fmt.Println(propt)
	 reader :=bufio.NewReader(os.Stdin)
	text,err :=reader.ReadString('\n')

	if err != nil{
		return ""
	}
	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r")
	return text
}