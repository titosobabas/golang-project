package databases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"taskstodo/crud"
)

type Local struct {
}

type tasksitems struct {
	items []crud.Task `json:"items"`
}

func (l *Local) GetItemsJsonFile() ([]crud.Task, error) {

	// Open our jsonFile
	jsonFile, err := os.Open("mydata.json")
	if err != nil {
		// fmt.Println(err)
	}

	// fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	// we initialize our Users array
	var tasksitems tasksitems

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &tasksitems.items)

	//fmt.Println("VAlores recuperados")
	//fmt.Println(tasksitems)
	return tasksitems.items, err

}

func (l Local) AddItemJsonFile(taskToSave crud.Task) (int, error) {
	//...................................
	//Writing struct type to a JSON file
	//...................................

	jsonItems, err := l.GetItemsJsonFile()
	if err != nil {
		return -1, err
	}

	itemsToSave := make([]crud.Task, 0)
	for i := 0; i < len(jsonItems); i++ {
		itemsToSave = append(itemsToSave, jsonItems[i])
	}
	if taskToSave.Id <= 0 {
		taskToSave.Id = len(itemsToSave) + 1
	}
	itemsToSave = append(itemsToSave, taskToSave)

	content, err := json.Marshal(itemsToSave)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("mydata.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return 1, nil
}

func (l *Local) GetAll() ([]crud.Task, error) {
	//TODO implement me
	tasks, err := l.GetItemsJsonFile()
	return tasks, err
}

func (l Local) UpdateItem(id int) (crud.Task, error) {
	taskDeleted, err := l.DeleteItem(id)
	if err != nil {
		return crud.Task{}, err
	}
	taskUpdated := crud.Task{
		Id:          taskDeleted.Id,
		Name:        taskDeleted.Name,
		Description: taskDeleted.Description,
		Timestamp:   taskDeleted.Timestamp,
		Done:        true,
	}
	_, err = l.AddItemJsonFile(taskUpdated)
	if err != nil {
		return crud.Task{}, err
	}
	return taskUpdated, nil
}

func (l Local) UpdateJsonFile(tasksToSave []crud.Task) (int, error) {
	content, err := json.Marshal(tasksToSave)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("mydata.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return 1, nil
}

func (l Local) DeleteItem(id int) (crud.Task, error) {
	// first we get all the tasks stored in the json file
	tasks, _ := l.GetItemsJsonFile()
	// iterating each json item until we find the one who match with the id sent
	newTasks := make([]crud.Task, 0)

	var taskFound crud.Task
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id != id {
			newTasks = append(newTasks, tasks[i])
		}
		if tasks[i].Id == id {
			taskFound = tasks[i]
		}
	}
	_, err := l.UpdateJsonFile(newTasks)
	return taskFound, err
}

func (l Local) CreateItem(task crud.Task) (crud.Task, error) {
	//TODO implement me
	taskToSave := task
	_, err := l.AddItemJsonFile(taskToSave)
	if err != nil {
		return taskToSave, err
	}
	return taskToSave, err
}

func (l Local) GetItem(id int) (crud.Task, error) {
	tasks, _ := l.GetItemsJsonFile()
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Id == id {
			return tasks[i], nil
		}
	}
	return crud.Task{}, fmt.Errorf("No item found with this id")
}

func NewLocal() *Local {
	return &Local{}
}
