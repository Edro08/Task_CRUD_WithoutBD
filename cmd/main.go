package main

import (
	"fmt"
	"garcia-eduardo-0323/internal/task/dominio"
	"garcia-eduardo-0323/internal/task/operation"
	"log"
)

func main() {

	var option string
	var option2 string
	var tasks []dominio.Task

	loadData, err := operation.NewLoadDataImpl().Load()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, v := range loadData {
		tasks = append(tasks, v)
	}

	ListData := operation.NewListDataImpl()
	AddData := operation.NewAddDataImpl()

	for option != "5" {
		Menu()

		fmt.Println("Digitar Opcion:")
		_, err := fmt.Scanln(&option)
		if err != nil {
			return
		}

		switch option {
		case "1":
			id, name, description, status, dateStart, dateFinally, priority := AddTask()
			task, err := AddData.Add(id, name, description, status, dateStart, dateFinally, priority)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			tasks = append(tasks, task)
		case "2":
			ListTaskMenu()

			fmt.Println("Digitar Opcion:")
			_, err := fmt.Scanln(&option2)
			if err != nil {
				return
			}

			err = ListData.List(option2, tasks)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		case "3":
			fmt.Println("two")
		case "4":
			fmt.Println("two")
		case "5":
			fmt.Println("Adios!")
		}
	}

}

func Menu() {
	//Inicio de pantalla
	fmt.Println("Task Management: ")
	fmt.Println("1- Add Task")
	fmt.Println("2- List Task")
	fmt.Println("3- Update Task")
	fmt.Println("4- Delete Task")
	fmt.Println("5- Salir")
}

func ListTaskMenu() {
	fmt.Println("	List Task:")
	fmt.Println("	1- All")
	fmt.Println("	2- Pendientes")
	fmt.Println("	3- Priority")
}

func AddTask() (id, name, description, status, dateStart, dateFinally, priority string) {
	var id2, name2, description2, status2, dateStart2, dateFinally2, priority2 string
	fmt.Println("Digitar Id:")
	_, err := fmt.Scanln(&id2)
	if err != nil {
		return
	}
	fmt.Println("Digitar Name:")
	_, err = fmt.Scanln(&name2)
	if err != nil {
		return
	}

	fmt.Println("Digitar Description:")
	_, err = fmt.Scanln(&description2)
	if err != nil {
		return
	}

	fmt.Println("Digitar Status:")
	_, err = fmt.Scanln(&status2)
	if err != nil {
		return
	}
	fmt.Println("Digitar Fecha Inicio:")
	_, err = fmt.Scanln(&dateStart2)
	if err != nil {
		return
	}

	fmt.Println("Digitar Fecha FInalizacion:")
	_, err = fmt.Scanln(&dateFinally2)
	if err != nil {
		return
	}

	fmt.Println("Digitar Prioridad:")
	_, err = fmt.Scanln(&priority2)
	if err != nil {
		return
	}

	return id2, name2, description2, status2, dateStart2, dateFinally2, priority2
}
