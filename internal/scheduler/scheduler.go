package scheduler

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

// var tm *TaskManager

func ping() {
	fmt.Println("pong")
}

type TaskManager struct {
	scheduler *gocron.Scheduler
	tasks     map[string]*Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		scheduler: gocron.NewScheduler(time.Local),
		tasks:     make(map[string]*Task),
	}
}
func (tm *TaskManager) CreateTask(id string, name string, interval time.Duration, fn func()) error {
	fmt.Printf("Creating task with ID %s and  %v\n", id, interval)
	if _, exists := tm.tasks[id]; exists {
		return fmt.Errorf("task with ID %s already exists", id)
	}

	job, err := tm.scheduler.Every(interval).Do(fn)
	if err != nil {
		return fmt.Errorf("failed to schedule task: %v interval %v   ", err, interval)
	}
	task := &Task{
		ID:       id,
		Name:     name,
		Interval: time.Duration(interval.Seconds()),
		Function: fn,
		IsActive: true,
		NextRun:  job.NextRun(),
	}

	tm.tasks[id] = task
	return nil
}

func (tm *TaskManager) UpdateTask(id string, newInterval time.Duration) error {
	task, exists := tm.tasks[id]
	if !exists {
		return fmt.Errorf("task with ID %s not found", id)
	}

	tm.scheduler.RemoveByTag(id)

	job, err := tm.scheduler.Every(newInterval).Do(task.Function)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}

	task.Interval = newInterval
	task.NextRun = job.NextRun()
	return nil
}

func (tm *TaskManager) DeleteTask(id string) error {
	if _, exists := tm.tasks[id]; !exists {
		return fmt.Errorf("task with ID %s not found", id)
	}

	tm.scheduler.RemoveByTag(id)
	delete(tm.tasks, id)
	return nil
}

func (tm *TaskManager) GetTask(id string) (*Task, error) {
	task, exists := tm.tasks[id]
	if !exists {
		return nil, fmt.Errorf("task with ID %s not found", id)
	}
	return task, nil
}

func (tm *TaskManager) Start() {
	tm.scheduler.StartAsync()
}

func (tm *TaskManager) Stop() {
	tm.scheduler.Stop()
}

func main() {
	// Example usage
	tm := NewTaskManager()

	// Create a sample task
	err := tm.CreateTask("task1", "Print Hello", 5*time.Second, func() {
		fmt.Println("Hello, World!")
	})
	if err != nil {
		fmt.Printf("Error creating task: %v\n", err)
		return
	}

	tm.Start()

	time.Sleep(10 * time.Second)
	err = tm.UpdateTask("task1", 2*time.Second)
	if err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return
	}

	time.Sleep(10 * time.Second)
	err = tm.DeleteTask("task1")
	if err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
		return
	}

	select {}
}
