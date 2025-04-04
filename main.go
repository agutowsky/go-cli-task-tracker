package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	TaskId      int    `json:"taskId"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	CompletedAt string `json:"completedAt"`
	IsDeleted   bool   `json:"isDeleted"`
}

func main() {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var tasks []Task

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Error: Missing argument.")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":

	case "update":

	case "delete":

	case "list":
		var pending []Task
		var active []Task
		var complete []Task

		for _, task := range tasks {
			switch task.Status {
			case "pending":
				pending = append(pending, task)
			case "active":
				active = append(active, task)
			case "complete":
				complete = append(complete, task)
			}
		}

		fmt.Printf("You have %d task(s)\n", len(tasks))
		fmt.Println("> Pending")
		for _, task := range pending {
			fmt.Printf("\t%d > \"%s\"\n", task.TaskId, task.Title)
		}

		fmt.Println("> Active")
		for _, task := range active {
			fmt.Printf("\t%d > \"%s\"\n", task.TaskId, task.Title)
		}

		fmt.Println("> Complete")
		for _, task := range complete {
			fmt.Printf("\t%d > \"%s\"\n", task.TaskId, task.Title)
		}

	default:
		fmt.Println("Unknown command:", cmd)
	}
}
