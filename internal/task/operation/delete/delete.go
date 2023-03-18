package delete

import (
	"garcia-eduardo-0323/internal/task/dominio"
)

type IDelete interface {
	Delete(id string, tasks []dominio.Task) (dominio.Task, error)
}

type DeleteData struct {
}

func NewDeleteDataImpl() *DeleteData {
	return &DeleteData{}
}

func (p *DeleteData) Delete(id string, tasks []dominio.Task) ([]dominio.Task, error) {
	var out []dominio.Task
	Id, err := dominio.NewTaskID(id)
	if err != nil {
		return tasks, err
	}

	for _, v := range tasks {
		if v.Id != Id {
			out = append(out, v)
		}
	}

	return out, nil
}
