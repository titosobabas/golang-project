package crud

import "testing"

func TestCrudCreateTask(t *testing.T) {
	tables := []struct {
		task Task
		m    string
	}{
		{
			Task{
				Name:        "task",
				Description: "task desc",
				Timestamp:   "xxxxxx",
				Done:        false,
			},
			"Task created successfully!",
		},
	}

	for _, table := range tables {
		_, err := CreateTask(table.task)
	}
}
