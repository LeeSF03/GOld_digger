package main

import (
	"GOld_digger/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is LeeSF's programming language!\nHave fun digging!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
