package main

import (
	"fmt"

	"github.com/BekirAktas/go-todo-app/models"
)

const (
	todoFile = "todos.json"
)

func main() {	
	fmt.Print("1 = add a new task")
	fmt.Print("2 = list of all tasks")
	fmt.Print("3 = delete a task")
	var userInput int;

	fmt.Scan(&userInput);

	switch userInput {
		case 1:
			var userInputContent string
			fmt.Println("Gorev icerigini yazin")
			fmt.Scan(&userInputContent)
			newTask := models.NewTask(userInputContent);
			newTask.StoreToJsonFile();
			fmt.Println("yeni bir gorev eklendi")
		case 2:
			fmt.Println("ikiye basildi")
		case 3:
			fmt.Println("uce basildi")
	}



	// jsonFile, err := os.Open("todos.json")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// byteValue, _ := ioutil.ReadAll(jsonFile)
	// var tasks models.Tasks
	// json.Unmarshal(byteValue, &tasks)

	// for i := 0; i < len(tasks.Tasks); i++ {
	// 	fmt.Println("Task Content: " + tasks.Tasks[i].GetContent())
	// 	fmt.Println("Is Task Done: " + tasks.Tasks[i].GetIsDone())
	// 	fmt.Println("Created At: " + tasks.Tasks[i].GetCreatedAt().String())
	// 	fmt.Println("Completed At: " + strconv.Itoa(tasks.Tasks[i].GetCompletedat().Hour()))
	// }
}