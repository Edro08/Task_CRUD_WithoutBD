package load

import (
	"garcia-eduardo-0323/internal/task/dominio"
)

type ILoad interface {
	Load([]dominio.Task) ([]dominio.Task, error)
}

type LoadData struct {
}

func NewLoadDataImpl() *LoadData {
	return &LoadData{}
}

func (p *LoadData) Load(tasks []dominio.Task) ([]dominio.Task, error) {

	task, err := dominio.NewTask("1", "Leer", "Leer un Libro", "1", "2023-03-08", "2023-03-08", "3")
	if err == nil {
		tasks = append(tasks, task)
	}

	task, err = dominio.NewTask("2", "Escuchar", "Escuchar Musica", "2", "2023-03-08", "2023-03-08", "2")
	if err == nil {
		tasks = append(tasks, task)
	}

	task, err = dominio.NewTask("3", "Trabajar", "Trabajar en la casa", "3", "2023-03-08", "2023-03-08", "1")
	if err == nil {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
