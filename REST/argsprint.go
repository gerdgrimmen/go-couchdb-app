package main

import (
	"fmt"
	"os"
)

// functionName is
var functionName string

func main() {

	if len(os.Args) == 1 { /* os.Args[0] is "compiler" or "compiler.exe" */
		//fmt.Println(os.Args, os.Args[0])
		help()
		return
	} else if len(os.Args) >= 2 {
		functionName = os.Args[1]
		fmt.Println(os.Args, os.Args[1], functionName)
		switch functionName {
		case "create":
			fmt.Println("create")
			createDB()
		case "select":
			fmt.Println("select")
			selectDB()
		case "save":
			fmt.Println("save")
			saveFile()
		case "all_db":
			fmt.Println("all_db")
			allDB()
		case "help":
			fmt.Println("help")
			help()
		default:
			fmt.Println("Functionname not found: ", functionName)
			os.Exit(3)
		}
	}
}

func createDB() {
	functionName = "asd"
	fmt.Println(functionName)
}

func selectDB() {

}

func saveFile() {

}

func allDB() {

}

func help() {
	fmt.Println("help: ")
}

func checkConfig() bool {
	return true
}
