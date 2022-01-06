package cli

import (
	"bufio"
	"fmt"
	"os"
	"taskstodo/crud"
	"time"
)

type Cli struct {
	Option          int
	Id              int
	Cm              crud.CrudManager
	TaskName        string
	TaskDescription string
}

func NewCli(cm crud.CrudManager) *Cli {
	return &Cli{Cm: cm}
}

func (c *Cli) DisplayHeaders(headerText string, instructions string) {
	fmt.Println("------------------------------------------")
	fmt.Println(fmt.Sprintf("%s", headerText))
	fmt.Println("------------------------------------------")
	if len(instructions) > 0 {
		fmt.Print(fmt.Sprintf("%s", instructions))
	}
}

/*func (c *Cli) ReturningToMainMenu(finalMessage string) {
	fmt.Println("")
	fmt.Println(finalMessage)
	fmt.Println("We'll take you back to the Main Menu in 2 seconds!")
	time.Sleep(time.Second * 2)

	c.ShowMenu()
}*/

func (c *Cli) ReturningToMainMenuType(message string) {
	fmt.Print(message)
	fmt.Scanln()
	c.ShowMenu()
}

func (c *Cli) ShowMessageWithWait(message string) {
	fmt.Println(message)
	time.Sleep(time.Second * 2)
}

func (c *Cli) DeleteTask() {

	c.DisplayHeaders("You're now in: Delete Task", "Hey there! Please give the id that you want to delete!: ")
	fmt.Scanf("%d", &c.Id)

	c.ShowMessageWithWait("Let me search it in our database and in case the Id exists, we're gonna updated it! Wait 2 seconds...")

	_, err := c.Cm.DeleteTask(c.Id)
	if err != nil {
		fmt.Printf("Oops! Something went wrong!: [%v]", err)
		return
	}

	c.ReturningToMainMenuType("Very well! We've delete your task successfully! Type anything to go back Main Menu:")

}

func (c *Cli) UpdateTask() {

	c.DisplayHeaders("You're now in: Update Task", "Hey there! Please type the id that you want to update as done!: ")
	fmt.Scanf("%d", &c.Id)

	c.ShowMessageWithWait("Very well!! Let me search it in our database and in case the Id exists, we're gonna updated it! Wait 2 seconds...")

	_, err := c.Cm.UpdateTask(c.Id)
	if err != nil {
		fmt.Printf("Oops! Something went wrong!: [%v]", err)
		return
	}

	c.ReturningToMainMenuType("Very well! We've updated your task as DONE! Type anything to go back Main Menu:")
}

func (c *Cli) CreateTask() {

	scanner := bufio.NewScanner(os.Stdin)

	c.DisplayHeaders("You're now in: Create Task", "")

	fmt.Print("Please give the task NAME: ")
	scanner.Scan()
	inputName := scanner.Text()
	c.TaskName = inputName

	fmt.Print("Please give the task DESCRIPTION: ")
	scanner.Scan()
	inputDesc := scanner.Text()
	c.TaskDescription = inputDesc

	task := c.Cm.NewEmptyTask(
		c.TaskName,
		c.TaskDescription,
	)
	msg, err := c.Cm.CreateTask(task)

	if err != nil {
		fmt.Println(msg)
		return
	}

	c.ReturningToMainMenuType("Very well! We've deleted your task successfully! Type anything to go back Main Menu:")
}

func (c *Cli) PrintTemplateTask(task crud.Task) {
	fmt.Println("\n**********************************")
	fmt.Printf("Id: %d \n", task.Id)
	fmt.Printf("Name: %s \n", task.Name)
	fmt.Printf("Description: %s \n", task.Description)
	fmt.Printf("Timestamp: %s \n", task.Timestamp)
	fmt.Printf("Done: %t \n", task.Done)
	fmt.Println("**********************************\n")
}

func (c *Cli) ShowAllTasks() {
	//tasks := make([]crud.Task, 0)

	c.DisplayHeaders("You're now in: Show All Tasks", "")

	tasks, err := c.Cm.GetTasks()

	if err != nil {
		fmt.Printf("Oops! Something went wrong!: [%v]", err)
		return
	}

	fmt.Printf("Tasks found: %d", len(tasks))
	fmt.Println("")
	if len(tasks) > 0 {
		for _, task := range tasks {
			c.PrintTemplateTask(task)
		}
	}

	c.ReturningToMainMenuType("Type anything to go back Main Menu: ")

}

func (c *Cli) RunAction() {
	switch c.Option {
	case 1:
		c.CreateTask()
	case 2:
		c.UpdateTask()
	case 3:
		c.DeleteTask()
	case 4:
		c.ShowAllTasks()
	}
}

func (c *Cli) ShowMenu() {

	fmt.Println("Task system. Please choose any option below:")
	fmt.Println("")

	fmt.Println("1. Create task")
	fmt.Println("2. Update task")
	fmt.Println("3. Delete task")
	fmt.Println("4. Show all tasks")
	// fmt.Println("5. Show a single task by id")

	fmt.Print("Enter your option: ")
	fmt.Scanln(&c.Option)

	c.RunAction()

}
