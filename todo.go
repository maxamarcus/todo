package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Represents a todo item.
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// Represents a list of todo items.
type List []item

// Add
// Add an item to the todo list.
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete
// Mark an item as complete.
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist.", i)
	}
	// Adjust index to 0-index.
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	return nil
}

// Delete
// Remove an item from the list.
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
    // Adjust index to 0-index.
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Save
// Save the todo list to a file.
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, js, 0644)
}

// Get
// Load a todo list from a file.
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

// String
// Prints a formatted list.
// Implements fmt.Stringer interface.
func (l *List) String() string {
    formatted := ""
    for k, t := range *l {
        prefix := "  "
        if t.Done {
            prefix = "X "
        }
        formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
    }
    return formatted
}
