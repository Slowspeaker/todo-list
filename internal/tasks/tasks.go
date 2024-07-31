package tasks

import (
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Task struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	ActiveAt time.Time `json:"activeAt"`
	Done     bool      `json:"done"`
}

var tasks = []Task{}
var mu sync.Mutex

func CreateTask(title string) (Task, error) {
	mu.Lock()
	defer mu.Unlock()

	if len(title) > 200 {
		return Task{}, errors.New("title too long")
	}

	newTask := Task{
		ID:       generateID(),
		Title:    title,
		ActiveAt: time.Now(),
		Done:     false,
	}
	tasks = append(tasks, newTask)
	return newTask, nil
}

func GetTasks(status string) []Task {
	mu.Lock()
	defer mu.Unlock()

	var filteredTasks []Task
	now := time.Now()

	for _, task := range tasks {
		if status == "active" && !task.Done && task.ActiveAt.Before(now) {
			if isWeekend(task.ActiveAt) {
				task.Title = "ВЫХОДНОЙ - " + task.Title
			}
			filteredTasks = append(filteredTasks, task)
		} else if status == "done" && task.Done {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks
}

func GetTaskByID(id string) (Task, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, errors.New("task not found")
}

func UpdateTask(id string, title string, activeAt time.Time) error {
	mu.Lock()
	defer mu.Unlock()

	if len(title) > 200 {
		return errors.New("title too long")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = title
			tasks[i].ActiveAt = activeAt
			return nil
		}
	}

	return errors.New("task not found")
}

func DeleteTask(id string) error {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("task not found")
}

func MarkTaskDone(id string) error {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			return nil
		}
	}

	return errors.New("task not found")
}

func generateID() string {
	return uuid.New().String()
}

func isWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
