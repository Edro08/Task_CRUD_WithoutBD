package list

import (
	"fmt"
	"garcia-eduardo-0323/internal/task/dominio"
	"sort"
)

type IList interface {
	List(option string, id string, tasks []dominio.Task) error
}

type ListData struct {
}

func NewListDataImpl() *ListData {
	return &ListData{}
}

func (p *ListData) List(option string, id string, tasks []dominio.Task) error {

	switch option {
	case "1":
		_ = p.All(tasks)
	case "2":
		_ = p.OrderByPriorityAndStatus(tasks)
	case "4":
		_ = p.onlyTask(id, tasks)
	}
	return nil
}

func (p *ListData) All(tasks []dominio.Task) error {

	for _, v := range tasks {
		fmt.Printf("Id: %v, Name: %s, Description: %v, Status: Pendiente, Date Start: %s, Date Finally: %v, Priority: %s \n", v.IDs(), v.Name, v.Description, v.DateStart, v.DateFinally, v.PriorityName(v.Prioritys()))
	}

	return nil
}

func (p *ListData) OrderByPriorityAndStatus(tasks []dominio.Task) error {

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Prioritys() < tasks[j].Prioritys()
	})

	for _, v := range tasks {
		if v.Statuses() == 1 {
			fmt.Printf("Id: %v, Name: %s, Description: %v, Status: Pendiente, Date Start: %s, Date Finally: %v, Priority: %s \n", v.IDs(), v.Name, v.Description, v.DateStart, v.DateFinally, v.PriorityName(v.Prioritys()))
		}
	}
	return nil
}

func (p *ListData) onlyTask(id string, tasks []dominio.Task) error {
	Id, err := dominio.NewTaskID(id)
	if err != nil {
		return err
	}

	for _, v := range tasks {
		if v.Id == Id {
			fmt.Println(v)
		}
	}
	return nil
}
