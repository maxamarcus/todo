package main

import (
    "flag"
	"fmt"
	"os"
	"todo"
)

const todoFileName = ".todo.json"

func main() {

    // -h option
    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(),
            "%s tool. By Max Marcus\n", os.Args[0])
        fmt.Fprintf(flag.CommandLine.Output(),
            "Taken from \"Powerful Command Line Applications in Go\"\n" +
            "by Ricardo Gerardi, Pragmatic Bookshelf (c) 2020.\n")
        fmt.Fprintf(flag.CommandLine.Output(),
            "Usage information:\n")
        flag.PrintDefaults()
    }

    // Command line flags.
    // Return values of these functions are **pointers.
    // parameters -- flag name, default value, help message
    task := flag.String("task", "", "Task to add to todo list.")
    list := flag.Bool("list", false, "List all tasks.")
    complete := flag.Int("complete", 0, "Item number to mark as complete.")
    flag.Parse()

    // Initialize item list.
	l := &todo.List{}

	// Load the todo list from a file.
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

    // Follow according to flags.
	switch {

    // Print the todo list.
    case *list:
        for _, item := range *l {
            if !item.Done {
                fmt.Println(item.Task)
            }
        }
    
    // Mark an item as complete.
    case *complete > 0:
        
        // Mark the item of the given number as complete.
        if err := l.Complete(*complete); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }

        // Save the list.
        if err := l.Save(todoFileName); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }

    // Add a task.
    case *task != "":
        l.Add(*task)

        // Save.
        if err := l.Save(todoFileName); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }

    // Invalid flag.
	default:
        fmt.Fprintln(os.Stderr, "Invalid option.")
        os.Exit(1)
    }
}
