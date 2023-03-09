package operation

import (
	"fmt"
	"garcia-eduardo-0323/internal/task/dominio"
)

type IDelete interface {
	Delete(id string) (dominio.Task, error)
}

type DeleteData struct {
}

func NewDeleteDataImpl() *DeleteData {
	return &DeleteData{}
}

func (p *DeleteData) Delete(id string) ([]dominio.Task, error) {

	fmt.Println("proximamente!")
	fmt.Println("")
	fmt.Println("")
	return []dominio.Task{}, nil
}
