package main 

import (
	"os"
	"fmt"
)

func main(){
	args := os.Args
	// fmt.Println(args)
	if len(args) < 2{
		fmt.Println("Need more args")
		os.Exit(1)
	}

	// var todolist []string
	todolist := []string{}
	switch args[1]{
	case "add":
		if len(args) < 3{
			fmt.Println("Need more args for add")
			os.Exit(1)
		}
		todolist = append(todolist, args[2] )
		fmt.Println("Added:", args[2])
	case "list":
		fmt.Println(len(todolist),"items will be listed")
		for _,item := range todolist{
			fmt.Println(item)
		}
		// for i:=0;i<len(todolist);i++{
		// 	fmt.Println(todolist[i])
		// }
		//handle list
	default: 
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}