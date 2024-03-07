package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID   int
	Name string
	Done bool
}

type TodoList struct {
	Todos []Todo
}

func AddTodo(name string, tl *TodoList) {
	id := len(tl.Todos) + 1
	task := Todo{ID: id, Name: name, Done: false}
	tl.Todos = append(tl.Todos, task)
	fmt.Printf("New Todo '%s' added with ID %d\n", name, id)
}

func DeleteTodo(id int, tl *TodoList) {
	for i, task := range tl.Todos {
		if task.ID == id { // check on how to remove from a slice
			tl.Todos = append(tl.Todos[:i], tl.Todos[i+1:]...)
			fmt.Printf("Todo with ID %d deleted\n", id)
			return
		}
	}
	fmt.Printf("Todo with ID %d not found\n", id)
}

func MarkAsDone(id int, tl *TodoList) {
	for i, task := range tl.Todos {
		if task.ID == id {
			tl.Todos[i].Done = true
			fmt.Printf("Todo has been marked as done")
			return
		}
	}
}

func ViewAll(tl *TodoList) {
	for _, todo := range tl.Todos {

		fmt.Printf("%d: %s , status = (%t)\n", todo.ID, todo.Name, todo.Done)
	}
}

func ViewAllDoneTodos(tl *TodoList) {

	for _, todo := range tl.Todos {
		if todo.Done {
			fmt.Printf("%d: %s \n", todo.ID, todo.Name)
		}
	}
}

func ViewAllNotDoneTodos(tl *TodoList) {

	for _, todo := range tl.Todos {
		if !todo.Done {
			fmt.Printf("%d: %s \n", todo.ID, todo.Name)
		}
	}
}

func main() {
	TodoList := TodoList{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add Todo")
		fmt.Println("2. Delete Todo")
		fmt.Println("3. Complete Todo")
		fmt.Println("4. View All Todos")
		fmt.Println("5. View All Done Todos")
		fmt.Println("6. View All Not Done Todos")
		fmt.Println("7. Exit")
		fmt.Print("Enter option number: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Print("Enter task name: ")
			scanner.Scan()
			taskName := scanner.Text()
			AddTodo(taskName, &TodoList)
		case "2":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			DeleteTodo(id, &TodoList)
		case "3":
			fmt.Print("Enter task ID to mark as complete: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			MarkAsDone(id, &TodoList)
		case "4":
			ViewAll(&TodoList)
		case "5":
			ViewAllDoneTodos(&TodoList)
		case "6":
			ViewAllNotDoneTodos(&TodoList)
		case "7":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
