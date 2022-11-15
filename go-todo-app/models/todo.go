package models

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

type Tasks struct {
	Tasks []task `json:"tasks"`
}
type task struct {
	content string
	isDone bool
	createdAt time.Time
	completedAt time.Time
}

func NewTask(content string) task{
	newTask :=  task {
		content: content,
		isDone: false,
		createdAt: time.Now(),
		completedAt: time.Time{},
	}

	return newTask
}

func (t *task) StoreToJsonFile() error {
	content := t.GetContent();
	isDone := t.GetIsDone();
	createdAt := t.GetCreatedAt();
	completedAt := t.GetCompletedat();
	file, _ := json.MarshalIndent(
		"Content : " + content + 
		"Is Done : " + isDone + 
		"Created at : " + createdAt.String() +
		"Completed At : " + completedAt.String(), 
		"", " ");

	return ioutil.WriteFile("todos.json", file, 0644)
}

func (t *task) completeTask() {
	t.isDone = true;
}

func (t *task)deleteTask() {
	t = nil
}

func (t *task) GetContent() string{
	return t.content;
}

func (t *task) GetIsDone() string{
	return strconv.FormatBool(t.isDone)
}

func (t *task) GetCreatedAt() time.Time {
	return t.createdAt;
}

func (t *task) GetCompletedat() time.Time {
	return t.completedAt;
}