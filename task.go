package main

import (
	"encoding/json"
	"os"
)

// Task struct defines the structure of a task
type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks []Task

// Load tasks from a JSON file
func loadTasks() error {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return err
	}
	return nil
}

// Save tasks to a JSON file
func saveTasks() error {
	file, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", file, 0644)
}
