package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	// fmt.Println(args)

	if len(args) == 0 {
		fmt.Println("You need to pass in a valid file name")
		os.Exit(1)
	}

	dat, err := ioutil.ReadFile(args[0])
	check(err)
	fmt.Print(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
