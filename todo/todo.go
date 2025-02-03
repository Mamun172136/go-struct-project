package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string
}

func (todo Todo) Display()  {
	fmt.Println(todo.Content)
}

func (todo Todo)Save() error{
	fileName :="todo.json"
	json,err:=json.Marshal(todo)
	if err != nil{
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func New(content string)(Todo,error){

	if content == ""{
		return Todo{}, errors.New("content is empty")
	}
	return Todo{
		Content:content,
	},nil
}