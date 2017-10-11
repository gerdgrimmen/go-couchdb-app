package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 { /* os.Args[0] is "compiler" or "compiler.exe" */
		fmt.Println(os.Args, os.Args[0])

		return
	} else if len(os.Args) == 2 {
		functionName := os.Args[1]
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
	} else if len(os.Args) > 2 { /* os.Args[0] is "main" or "main.exe" */
		fmt.Println(os.Args, os.Args[1], os.Args[2])
		return
	}
}

func createDB() {

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
