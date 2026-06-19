package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func markTodos(filename string,index int){
	todos := loadTodos(filename)
	target := todos[index-1]
	if strings.HasPrefix(target,"[Done]"){
		fmt.Println("Already marked done")
		return
	} 
	todos[index-1]="[Done] " + target
	saveTodos(filename,todos)
}

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
	txt_file := "todos.txt"
	args := os.Args
	// fmt.Println(args)
	if len(args) < 2{
		fmt.Println("Need more args")
		os.Exit(1)
	}

	todolist := loadTodos(txt_file)
	switch args[1]{
	case "add":
		if len(args) < 3{
			fmt.Println("Need more args for add")
			os.Exit(1)
		}
		todolist = append(todolist, args[2] )
		fmt.Println("Added:", args[2])
		saveTodos(txt_file, todolist)
	case "list":
		fmt.Println(len(todolist),"items will be listed")
		for ind,item := range todolist{
			fmt.Printf("#%d. %s\n",ind+1,item)
		}
	case "done":
		if len(args) < 3{
			fmt.Println("Need index for marking done")
			os.Exit(1)
		}
		ind, err := strconv.Atoi(args[2])
		if err != nil || ind <=0 || ind > len(todolist){
			fmt.Println("Invalid index")
			os.Exit(1)
		}
		markTodos(txt_file,ind)
	default: 
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}