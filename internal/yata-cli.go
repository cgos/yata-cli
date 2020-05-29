//Package internal private application and library code
package internal

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/cgos/yata-cli/model"
)

// Yata Structure for yata application
type Yata struct {
	todolist  []*model.Todo
	fileStore string // os.UserHomeDir + .yata
}

// Read all TodoLists from a file.
// Pointers vs Values param/return: https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values/23551970#23551970
func (y *Yata) Read() []*model.Todo {
	// file, err := os.Open(y.fileStore)
	// if err != nil {
	// 	fmt.Println("error opening: ", err)
	// 	panic("Error while opening the file")
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(bufio.NewReader(file))

	// var todos []model.Todo
	// for scanner.Scan() {
	// 	todoName := scanner.Text()
	// 	item := model.Todo{Title: todoName, Done: false}
	// 	todos = append(todos, item)
	// }
	// todolist := TodoList{Name: "Main", Todo: todos}
	return nil
}

// Write all todos to file.
func (y *Yata) Write(todolist []model.Todo) {
	file, err := os.Create(y.fileStore)
	if err != nil {
		panic("Cannot open file")
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	data, err := json.Marshal(todolist)
	if err != nil {
		panic("Unable to marshal todolist")
	}

	nn, err := w.Write(data)
	if (nn < len(data)) || err != nill {
		panic("Error while writing to file")
	}
}

// Display a list of todos
func (y *Yata) Display() {
	y.todolist = y.Read()
	// fmt.Printf("%s \n", y.todolist.Name)
	// for j, item := range y.todolist.Todo {
	// 	fmt.Printf("\t%d: %s\n", j, item.Title)
	// }
}
