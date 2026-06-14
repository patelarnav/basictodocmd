package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadTodos(filename string) []string{
	//open file
	file,err := os.Open(filename)
	if err != nil{
		if os.IsNotExist(err){
			return []string{}
		}
		fmt.Println("error:",err)
		return nil
	}

	defer file.Close()

	todos := []string{}

	scanner := bufio.NewScanner(file)
	// read line by line by bfio scanner
	// loop statement returns false when it is unable to read line
	// whether by error or end
	for scanner.Scan() {
		todos = append(todos, scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		fmt.Println("scan error:", err)
		return nil
	}

	return todos
}

func saveTodos(filename string, todolist []string) {

	// 	//open/ create file
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil{
		fmt.Println("error:",err)
		os.Exit(1)
	}
	defer file.Close()
	// fmt.Println("file exists")

	writer := bufio.NewWriter(file)
	//write each todo as a line
	for _, todo := range todolist{
		fmt.Fprintln(writer,todo)
	}
	writer.Flush()

}

func main(){
	args := os.Args
	// fmt.Println(args)
	if len(args) < 2{
		fmt.Println("Need more args")
		os.Exit(1)
	}

	todolist := loadTodos("todos.txt")
	switch args[1]{
	case "add":
		if len(args) < 3{
			fmt.Println("Need more args for add")
			os.Exit(1)
		}
		todolist = append(todolist, args[2] )
		fmt.Println("Added:", args[2])
		saveTodos("todos.txt", todolist)
	case "list":
		fmt.Println(len(todolist),"items will be listed")
		for _,item := range todolist{
			fmt.Println(item)
		}
	default: 
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}