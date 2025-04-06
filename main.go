package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)



func main(){
	command:=flag.String("command", "command", "command",)
	run:=flag.String("run", "init", "init",)
	flag.Parse();
	fmt.Println(*run)
	switch *command {
	case "command":
		fmt.Println("Check for help with -command=help")
	case "start":
		os.MkdirAll("mongo-migrations", os.ModePerm)
		os.Create("mongo-migrations/main.go")
	case "main":
		exec.Command("go run mongo-migrations/main.go")
	case "create-migration-file":
		//create file with current date including seconds
		fileName:="mongo-migrations/"+time.Now().Format("2006-01-02-15-04-05")+".go"
		exec.Command("go run " + fileName)
	case "help":
		fmt.Println("To create migration folder run -command=start")
		fmt.Println("To create default file run -command=main")
		fmt.Println("To create migration file run -command=create-migration-file")
	default:
		fmt.Println("Unknown command")
		
	}
	
}