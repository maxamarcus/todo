package main

import (
	"fmt"
	"os"
	"strings"
	"todo"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	// Load the todo list from a file.
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {

	// No args -- print list
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}

		// Concatenate args and add to list.
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
