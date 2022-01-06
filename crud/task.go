package crud

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
	Done        bool   `json:"done"`
}

func NewTask(id int, name string, description string, timestamp string, done bool) *Task {
	return &Task{Id: id, Name: name, Description: description, Timestamp: timestamp, Done: done}
}

type TaskSource interface {
	GetAll() ([]Task, error)
	GetItem(id int) (Task, error)
	UpdateItem(id int) (Task, error)
	DeleteItem(id int) (Task, error)
	CreateItem(task Task) (Task, error)
}
