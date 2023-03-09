package operation

import (
	"fmt"
	"garcia-eduardo-0323/internal/task/dominio"
)

type IUpdate interface {
	Update(id, nombre, description, status, dateStart, dateFinally, priority string) (dominio.Task, error)
}

type UpdateData struct {
}

func NewUpdateDataImpl() *UpdateData {
	return &UpdateData{}
}

func (p *UpdateData) Update(id, nombre, description, status, dateStart, dateFinally, priority string) ([]dominio.Task, error) {

	fmt.Println("proximamente!")
	fmt.Println("")
	fmt.Println("")
	return []dominio.Task{}, nil
}
