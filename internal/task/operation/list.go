package operation

import (
	"fmt"
	"garcia-eduardo-0323/internal/task/dominio"
)

type IList interface {
	List(option string, tasks []dominio.Task) error
}

type ListData struct {
}

func NewListDataImpl() *ListData {
	return &ListData{}
}

func (p *ListData) List(option string, tasks []dominio.Task) error {

	switch option {
	case "1":
		_ = p.All(tasks)
	case "2":
		_ = p.Pendientes(tasks)
	case "3":
		fmt.Println("three")

	}
	return nil
}

func (p *ListData) All(tasks []dominio.Task) error {

	for _, v := range tasks {
		fmt.Println(v)
	}

	return nil
}

func (p *ListData) Pendientes(tasks []dominio.Task) error {

	for _, v := range tasks {
		if v.Statuses() == 1 {
			fmt.Println(v)
		}
	}
	return nil
}
