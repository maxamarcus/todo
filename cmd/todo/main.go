package main

import (
    "bufio"
    "io"
    "strings"
    "flag"
	"fmt"
	"os"
	"todo"
)

// File name default.
var todoFileName = ".todo.json"


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
    add := flag.Bool("add", false, "Add task to todo list.")
    list := flag.Bool("list", false, "List all tasks.")
    complete := flag.Int("complete", 0, "Item number to mark as complete.")
    flag.Parse()

    // Initialize item list.
	l := &todo.List{}

    // Check if user defined env var for custom file name.
    if os.Getenv("TODO_FILENAME") != "" {
        todoFileName = os.Getenv("TODO_FILENAME")
    }

	// Load the todo list from the file.
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

    // Follow according to flags.
	switch {

    // Print the todo list.
    case *list:
        fmt.Print(l)
    
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
    case *add:
        t, err := getTask(os.Stdin, flag.Args()...)
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
        l.Add(t)

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


func getTask(r io.Reader, args ...string) (string, error) {

    // Command line arguments.
    if len(args) > 0 {
        return strings.Join(args, " "), nil
    }

    // Standard input arguments.
    s := bufio.NewScanner(r)
    s.Scan()
    if err := s.Err(); err != nil {
        return "", err
    }
    if len(s.Text()) == 0 {
        return "", fmt.Errorf("Task cannot be blank.")
    }
    return s.Text(), nil
}
