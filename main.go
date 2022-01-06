package main

import (
	"taskstodo/crud"
	"taskstodo/databases"
	"taskstodo/pkg/cli"
)

func main() {
	source := databases.NewLocal()

	initCrud := crud.NewCrud(source)

	initCli := cli.NewCli(initCrud)

	initCli.ShowMenu()

	/*task := initCrud.NewEmptyTask(
		"Hello there!",
		"I'm the best task ever!!!!",
	)
	fmt.Println(initCrud.CreateTask(task))*/

	// fmt.Println(initCrud.GetTasks())

	// fmt.Println(initCrud.DeleteTask(4))

	// fmt.Println(initCrud.UpdateTask(1))

	// fmt.Println(initCrud.GetTask(5))

}
