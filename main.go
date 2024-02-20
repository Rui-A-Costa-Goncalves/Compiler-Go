package main

import (
	"fmt"
	"goCompiler/repl"
	"os"
	"os/user"
)

func main () {
	user, errUser := user.Current()
	if errUser != nil {
		panic(errUser)
	}
	fmt.Println("Welcome to the compiler user",user.Username, "!")
	repl.Start(os.Stdin,os.Stdout)
}