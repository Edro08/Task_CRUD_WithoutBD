package add

import (
	"fmt"
	"garcia-eduardo-0323/internal/task/dominio"
)

type IAdd interface {
	Add(id, nombre, description, status, dateStart, dateFinally, priority string, tasks []dominio.Task) ([]dominio.Task, error)
}

type AddData struct {
}

func NewAddDataImpl() *AddData {
	return &AddData{}
}

func (p *AddData) Add(id, nombre, description, status, dateStart, dateFinally, priority string, tasks []dominio.Task) ([]dominio.Task, error) {

	Id, err := dominio.NewTaskID(id)
	if err != nil {
		return tasks, err
	}

	for _, v := range tasks {
		if v.Id == Id {
			return tasks, fmt.Errorf("id already exists")
		}
	}

	task, err := dominio.NewTask(id, nombre, description, status, dateStart, dateFinally, priority)
	if err != nil {
		return tasks, err
	}

	tasks = append(tasks, task)
	return tasks, nil
}
