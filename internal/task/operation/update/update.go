package update

import (
	"garcia-eduardo-0323/internal/task/dominio"
)

type IUpdate interface {
	Update(id, nombre, description, status, dateStart, dateFinally, priority string, tasks []dominio.Task) ([]dominio.Task, error)
}

type UpdateData struct {
}

func NewUpdateDataImpl() *UpdateData {
	return &UpdateData{}
}

func (p *UpdateData) Update(id, nombre, description, status, dateStart, dateFinally, priority string, tasks []dominio.Task) ([]dominio.Task, error) {
	var out []dominio.Task
	Id, err := dominio.NewTaskID(id)
	if err != nil {
		return tasks, err
	}

	for _, v := range tasks {
		if v.Id == Id {
			task, err := dominio.NewTask(id, nombre, description, status, dateStart, dateFinally, priority)
			if err != nil {
				return tasks, err
			}
			out = append(out, task)
		} else {
			out = append(out, v)
		}
	}

	return out, nil
}
