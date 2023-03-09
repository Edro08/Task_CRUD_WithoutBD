package operation

import (
	"garcia-eduardo-0323/internal/task/dominio"
)

type IAdd interface {
	Add(id, nombre, description, status, dateStart, dateFinally, priority string) (dominio.Task, error)
}

type AddData struct {
}

func NewAddDataImpl() *AddData {
	return &AddData{}
}

func (p *AddData) Add(id, nombre, description, status, dateStart, dateFinally, priority string) (dominio.Task, error) {

	task, err := dominio.NewTask(id, nombre, description, status, dateStart, dateFinally, priority)
	if err != nil {
		return dominio.Task{}, err
	}

	return task, nil
}
