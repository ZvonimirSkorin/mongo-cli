package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)



func main(){
	init:=flag.String("init", "init", "init",)
	run:=flag.String("run", "init", "init",)
	flag.Parse();
	fmt.Println(*run)
	if(*init == "start"){
		os.MkdirAll("mongo-migrations", os.ModePerm)
		os.Create("mongo-migrations/main.go")
	}
	 if(*run == "main"){
		exec.Command("go run mongo-migrations/main.go")
	 }
	
}