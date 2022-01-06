package crud

import (
	"github.com/itchyny/timefmt-go"
	"time"
)

type Crud struct {
	ts TaskSource
}

type CrudManager interface {
	GetTasks() ([]Task, error)
	CreateTask(task Task) (string, error)
	DeleteTask(id int) (string, error)
	UpdateTask(id int) (string, error)
	GetTask(id int) (Task, error)
	NewEmptyTask(name string, description string) Task
}

func NewCrud(ts TaskSource) *Crud {
	return &Crud{ts: ts}
}

func (c *Crud) CreateTimestamp() string {
	t := time.Now()
	t.Format("20060102150405")
	str := timefmt.Format(t, "%Y-%m-%d %H:%M:%S")
	return str
}

func (c *Crud) NewEmptyTask(name string, description string) Task {
	return Task{
		Name:        name,
		Description: description,
		Timestamp:   c.CreateTimestamp(),
		Done:        false,
	}
}

func (c *Crud) GetTasks() ([]Task, error) {
	items, err := c.ts.GetAll()
	return items, err
}

func (c *Crud) CreateTask(task Task) (string, error) {
	taskToSave := task
	_, err := c.ts.CreateItem(taskToSave)
	if err == nil {
		return "Task created successfully!", nil
	}
	return "Sorry, an error ocurrred!", err
}

func (c *Crud) DeleteTask(id int) (string, error) {
	taskId := id
	_, err := c.ts.DeleteItem(taskId)
	if err != nil {
		return "Something failed!", err
	}
	return "Task deleted!", nil
}

func (c *Crud) UpdateTask(id int) (string, error) {
	_, err := c.ts.UpdateItem(id)
	if err != nil {
		return "Something failed!", err
	}
	return "Task updated!", nil
}

func (c *Crud) GetTask(id int) (Task, error) {
	task, err := c.ts.GetItem(id)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
