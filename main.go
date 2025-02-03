package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"

	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("todo title")
	todo,err := todo.New(todoText)

	if err!= nil{
		fmt.Println(err)
	 }
	 err = todo.Save()
	 if err != nil{
		fmt.Println(err)
		return
	 }

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
 fmt.Println("note file save")
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