package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Martin-Martinez4/compiler_in_go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, this is the Monkey programming language\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
